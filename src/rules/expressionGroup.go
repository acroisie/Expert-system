package rules

import (
	"errors"
	"expert-system/src/v"
	"expert-system/src/factManager"
	"fmt"
)

var ExpressionGroupDisplayLogs bool = false

type ExpressionGroup struct {
    Op LogicalOperator
	LeftVariable *Variable
	RightVariable *Variable
	LeftExpressionGroup *ExpressionGroup
	RightExpressionGroup *ExpressionGroup
}

func (ep ExpressionGroup) solving() (v.Value, error) {
	LogEp(fmt.Sprintf("ExpressionGroup solving : %s", ep))

	if !ep.Op.isValid() {
		return v.UNDETERMINED, errors.New(fmt.Sprintf("Unknown operator: %s", ep.Op))
	}

	leftValue, err := solvingSide(ep.LeftVariable, ep.LeftExpressionGroup)
	if err != nil {
		return v.UNDETERMINED, err
	}
	rightValue, err := solvingSide(ep.RightVariable, ep.RightExpressionGroup)
	if err != nil {
		return v.UNDETERMINED, err
	}
	LogEp(fmt.Sprintf("%s : %s, %s : %s", ep.DisplaySide(LEFT), leftValue, ep.DisplaySide(RIGHT), rightValue))
	
	result, err := ep.Op.solve(leftValue, rightValue)
	if err != nil {
		return v.UNDETERMINED, err
	}
	LogEp(fmt.Sprintf("Solving : %s = %s", ep, result))

	return result, nil
}

func solvingSide(variable *Variable, expressionGroup *ExpressionGroup) (v.Value, error) {
	if variable != nil {
		fact, err := factManager.GetFactReferenceByLetter(variable.Letter)
		if err != nil {
			return v.UNDETERMINED, err
		}
		if variable.Not {
			return fact.Value.NOT(), nil
		}
		return fact.Value, nil
	} else if expressionGroup != nil {
		return expressionGroup.solving()
	} else {
		return v.UNDETERMINED, errors.New("Left side is empty")
	}
}

func (ep ExpressionGroup) deduction(result v.Value) error {
	LogEp(fmt.Sprintf("%s deduction with result : %s", ep, result))
	if result.Real() {
		leftValue, err := solvingSide(ep.LeftVariable, ep.LeftExpressionGroup)
		if err != nil {
			return err
		}
		rightValue, err := solvingSide(ep.RightVariable, ep.RightExpressionGroup)
		if err != nil {
			return err
		}

		if leftValue == v.UNKNOWN && rightValue.Real() {
			return ep.findOneUnknown(result, rightValue, LEFT)
		} else if leftValue.Real() && rightValue == v.UNKNOWN {
			return ep.findOneUnknown(result, leftValue, RIGHT)
		} else if leftValue == v.UNKNOWN && rightValue == v.UNKNOWN {
			return ep.findTwoUnknow(result)
		}
	}
	return nil
}

func sideDeduction(variable *Variable, expressionGroup *ExpressionGroup, newValue v.Value) error {
	if variable != nil {
		if variable.Not {
			newValue = newValue.NOT()
		}
		return factManager.SetFactValueByLetter(variable.Letter, newValue)
	} else {
		return expressionGroup.deduction(newValue)
	}
}

func (ep ExpressionGroup) findOneUnknown(res v.Value, know v.Value, side Side) error {
	var newValue v.Value
	if ep.Op == OR {
		newValue = res.FindUnknown_OR(know)
	} else if ep.Op == AND {
		newValue = res.FindUnknown_AND(know)
	} else {
		newValue = res.FindUnknown_XOR(know)
	}

	if side == LEFT {
		return sideDeduction(ep.LeftVariable, ep.LeftExpressionGroup, newValue)
	} else {
		return sideDeduction(ep.RightVariable, ep.RightExpressionGroup, newValue)
	}
}

func (ep ExpressionGroup) findTwoUnknow(res v.Value) error {
	var newLeftValue v.Value
	var newRightValue v.Value
	if ep.Op == OR {
		newLeftValue, newRightValue = res.FindTwoUnknown_OR()
	} else if ep.Op == AND {
		newLeftValue, newRightValue = res.FindTwoUnknown_AND()
	} else {
		newLeftValue, newRightValue = res.FindTwoUnknown_XOR()
	}

	if newLeftValue != v.UNKNOWN {
		return sideDeduction(ep.LeftVariable, ep.LeftExpressionGroup, newLeftValue)
	} else if newRightValue != v.UNKNOWN {
		return sideDeduction(ep.RightVariable, ep.RightExpressionGroup, newRightValue)
	}
	return nil
}

// DISPLAY

func LogEp(msg string) {
	if ExpressionGroupDisplayLogs {
		fmt.Println(fmt.Sprintf("ExpressionGroup - %s", msg))
	}
}

func (ep ExpressionGroup) DisplaySide(side Side) string {
	if side == LEFT {
		if ep.LeftVariable != nil {
			return ep.LeftVariable.String()
		} else {
			return ep.LeftExpressionGroup.String()
		}
	} else {
		if ep.RightVariable != nil {
			return ep.RightVariable.String()
		} else {
			return ep.RightExpressionGroup.String()
		}
	}
}

func (ep ExpressionGroup) String() string {
	return fmt.Sprintf("%s %s %s", ep.DisplaySide(LEFT), ep.Op, ep.DisplaySide(RIGHT))
}


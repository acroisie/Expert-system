package rules

import (
	"fmt"
	"errors"
)

type ExpressionGroup struct {
    Op LogicalOperator
	LeftVariable *Variable
	RightVariable *Variable
	LeftExpressionGroup *ExpressionGroup
	RightExpressionGroup *ExpressionGroup
}

func (ep ExpressionGroup) solving(facts []Fact, display bool) (Value, error) {
	if display {
		fmt.Println("ExpressionGroup solving : ", ep)
	}

	if !ep.Op.isValid() {
		return v.UNDETERMINED, errors.New(fmt.Sprintf("Unknown operator: %s", ep.Op))
	}

	leftValue, err := ep.solvingSide(true, facts)
	if err != nil {
		return v.UNDETERMINED, err
	}
	rightValue, err := ep.solvingSide(false, facts)
	if err != nil {
		return v.UNDETERMINED, err
	}
	if display {
		fmt.Println(fmt.Sprintf("ExpressionGroup solving leftValue %s : %s", ep.SideDisplay(true), leftValue))
		fmt.Println(fmt.Sprintf("ExpressionGroup solving rightValue %s : %s", ep.SideDisplay(false), rightValue))
	}
	
	result, err := ep.Op.solve(leftValue, rightValue, display)
	if err != nil {
		return v.UNDETERMINED, err
	}
	if display {
		fmt.Println(fmt.Sprintf("ExpressionGroup solving result: %s = %s", ep, result))
	}
	return result, nil
}

func (ep ExpressionGroup) solvingSide(side bool, facts []Fact) (Value, error) {
	if side {
		if ep.LeftVariable != nil {
			return ep.LeftVariable.GetValueByFacts(facts)
			return fac
		} else if ep.LeftExpressionGroup != nil {
			return ep.LeftExpressionGroup.solving(facts, false)
		} else {
			return v.UNDETERMINED, errors.New("Left side is empty")
		}
	} else {
		if ep.RightVariable != nil {
			return ep.RightVariable.GetValueByFacts(facts)
		} else if ep.RightExpressionGroup != nil {
			return ep.RightExpressionGroup.solving(facts, false)
		} else {
			return v.UNDETERMINED, errors.New("Right side is empty")
		}
	}
}

func (ep ExpressionGroup) findTwoUnknow(res Value, facts []Fact, change int) (int, error) {
	var newLeftValue Value
	var newRightValue Value
	if ep.Op == OR {
		newLeftValue, newRightValue = res.findTwoUnknown_OR()
	} else if ep.Op == AND {
		newLeftValue, newRightValue = res.findTwoUnknown_AND()
	} else {
		newLeftValue, newRightValue = res.findTwoUnknown_XOR()
	}

	if newLeftValue != v.UNKNOW {
		if ep.LeftVariable != nil{
			err := SetFactValueByLetter(facts, *ep.LeftVariable, newLeftValue, false)
			if err != nil {
				return change, err
			}
			return change + 1, nil
		} else {
			res, err := ep.LeftExpressionGroup.resultDeduction(newLeftValue, change, facts)
			return res, err
		}
	} else if newRightValue != v.UNKNOW {
		if ep.RightVariable != nil {
			err := SetFactValueByLetter(facts, *ep.RightVariable, newRightValue, false)
			if err != nil {
				return change, err
			}
			return change + 1, nil
		} else {
			res, err := ep.RightExpressionGroup.resultDeduction(newRightValue, change, facts)
			return res, err
		}
	}
	return change, nil
}

func (ep ExpressionGroup) findUnknow(res Value, know Value, side bool, facts []Fact, change int) (int, error) {
	var newValue Value
	if ep.Op == OR {
		newValue = res.findUnknown_OR(know)
	} else if ep.Op == AND {
		newValue = res.findUnknown_AND(know)
	} else {
		newValue = res.findUnknown_XOR(know)
	}

	if side {
		if ep.LeftVariable != nil {
			err := SetFactValueByLetter(facts, *ep.LeftVariable, newValue, false)
			if err != nil {
				return change, err
			}
			return change + 1, nil
		} else {
			res, err := ep.LeftExpressionGroup.resultDeduction(newValue, change, facts)
			return res, err
		}
	} else {
		if ep.RightVariable != nil {
			err := SetFactValueByLetter(facts, *ep.RightVariable, newValue, false)
			if err != nil {
				return change, err
			}
			return change + 1, nil
		} else {
			res, err := ep.RightExpressionGroup.resultDeduction(newValue, change, facts)
			return res, err
		}
	}
}

func (ep ExpressionGroup) resultDeduction(result Value, change int, facts []Fact) (int, error) {
	leftValue, err := ep.solvingSide(true, facts)
	if err != nil {
		return change, err
	}
	rightValue, err := ep.solvingSide(false, facts)
	if err != nil {
		return change, err
	}
	if result.Real() {
		if leftValue == v.UNKNOW && rightValue.Real() {
			change, err := ep.findUnknow(result, rightValue, true, facts, change)
			if err != nil {
				return change, err
			}
		} else if leftValue.Real() && rightValue == v.UNKNOW {
			change, err := ep.findUnknow(result, leftValue, false, facts, change)
			if err != nil {
				return change, err
			}
		} else if leftValue == v.UNKNOW && rightValue == v.UNKNOW {
			change, err := ep.findTwoUnknow(result, facts, change)
			if err != nil {
				return change, err
			}
		}
	}
	return change, nil
}

// DISPLAY

func (ep ExpressionGroup) String() string {
	return fmt.Sprintf("%s %s %s", ep.SideDisplay(true), ep.Op, ep.SideDisplay(false))
}

func (ep ExpressionGroup) SideDisplay(side ) string {
	if side {
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

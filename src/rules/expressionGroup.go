package rules

import (
	"expert-system/src/factManager"
	"expert-system/src/v"
	"fmt"
)

var ExpressionGroupDisplayLogs bool = false

type ExpressionGroup struct {
	Op                   LogicalOperator
	LeftVariable         *Variable
	RightVariable        *Variable
	LeftExpressionGroup  *ExpressionGroup
	RightExpressionGroup *ExpressionGroup
}

func (ep ExpressionGroup) solving() (v.Value, *v.Error) {
	// LogEp(fmt.Sprintf("ExpressionGroup solving : %s", ep))
	// LogEp(fmt.Sprintf("LeftVariable : %s", ep.LeftVariable))
	// LogEp(fmt.Sprintf("RightVariable : %s", ep.RightVariable))
	// LogEp(fmt.Sprintf("LeftExpressionGroup : %s", ep.LeftExpressionGroup))
	// LogEp(fmt.Sprintf("RightExpressionGroup : %s", ep.RightExpressionGroup))

	if !ep.Op.isValid() {
		return v.UNDETERMINED, &v.Error{Type: v.SOLVING, Message: fmt.Sprintf("Invalid operator : %s", ep.Op)}
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

func solvingSide(variable *Variable, expressionGroup *ExpressionGroup) (v.Value, *v.Error) {
	LogRule(fmt.Sprintf("Solving side : %s, %s", variable, expressionGroup))
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
		return v.UNDETERMINED, &v.Error{Type: v.SOLVING, Message: "Variable and ExpressionGroup are nil"}
	}
}

func (ep ExpressionGroup) deduction(result v.Value) *v.Error {
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
		} else if leftValue.Real() && rightValue.Real() {
			res, err := ep.Op.solve(leftValue, rightValue)
			if err != nil {
				return err
			}
			if res != result {
				return &v.Error{Type: v.CONTRADICTION, Message: fmt.Sprintf("in deduction: %s %s %s, for %s", leftValue, ep.Op, rightValue, ep)}
			}
		}
	}
	return nil
}

func sideDeduction(variable *Variable, expressionGroup *ExpressionGroup, newValue v.Value) *v.Error {
	if variable != nil {
		LogReasoning(fmt.Sprintf("so %s = %s\n", variable, newValue))
		if variable.Not {
			LogReasoning(fmt.Sprintf("%s = %s, so %s = %s\n", variable, newValue, variable.Letter, newValue.NOT()))
			newValue = newValue.NOT()
		}
		return factManager.SetFactValueByLetter(variable.Letter, newValue, false)
	} else {
		LogReasoning(fmt.Sprintf("so %s = %s\n", expressionGroup, newValue))
		return expressionGroup.deduction(newValue)
	}
}

func (ep ExpressionGroup) findOneUnknown(res v.Value, know v.Value, side Side) *v.Error {
	var newValue v.Value
	if ep.Op == OR {
		newValue = res.FindUnknown_OR(know)
	} else if ep.Op == AND {
		newValue = res.FindUnknown_AND(know)
	} else {
		newValue = res.FindUnknown_XOR(know)
	}

	LogReasoning(fmt.Sprintf("%s = %s, ", ep, res))
	if side == LEFT {
		return sideDeduction(ep.LeftVariable, ep.LeftExpressionGroup, newValue)
	} else {
		return sideDeduction(ep.RightVariable, ep.RightExpressionGroup, newValue)
	}
}

func (ep ExpressionGroup) findTwoUnknow(res v.Value) *v.Error {
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
		LogReasoning(fmt.Sprintf("%s = %s, ", ep, res))
		return sideDeduction(ep.LeftVariable, ep.LeftExpressionGroup, newLeftValue)
	} else if newRightValue != v.UNKNOWN {
		LogReasoning(fmt.Sprintf("%s = %s, ", ep, res))
		return sideDeduction(ep.RightVariable, ep.RightExpressionGroup, newRightValue)
	}
	return nil
}

func (ep ExpressionGroup) getFactOccurences(factListOccurence *map[rune]int) {
	if ep.LeftVariable != nil {
		(*factListOccurence)[ep.LeftVariable.Letter]++
	} else {
		ep.LeftExpressionGroup.getFactOccurences(factListOccurence)
	}
	if ep.RightVariable != nil {
		(*factListOccurence)[ep.RightVariable.Letter]++
	} else {
		ep.RightExpressionGroup.getFactOccurences(factListOccurence)
	}
}

func (ep ExpressionGroup) GetLetters() map[rune]struct{} {
	letters := make(map[rune]struct{})
	if ep.LeftVariable != nil {
		letters[ep.LeftVariable.Letter] = struct{}{}
	} else {
		for letter := range ep.LeftExpressionGroup.GetLetters() {
			letters[letter] = struct{}{}
		}
	}
	if ep.RightVariable != nil {
		letters[ep.RightVariable.Letter] = struct{}{}
	} else {
		for letter := range ep.RightExpressionGroup.GetLetters() {
			letters[letter] = struct{}{}
		}
	}
	return letters
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
	if ep.Op == NOTHING {
		return ep.DisplaySide(LEFT)
	}
	return fmt.Sprintf("%s %s %s", ep.DisplaySide(LEFT), ep.Op, ep.DisplaySide(RIGHT))
}

func (eg *ExpressionGroup) PrintAST(prefix string, isLast bool) string {
	var s string

	if eg.Op == NOTHING {
		if eg.LeftVariable != nil {
			if isLast {
				s += fmt.Sprintf("%s└── %s\n", prefix, eg.LeftVariable)
			} else {
				s += fmt.Sprintf("%s├── %s\n", prefix, eg.LeftVariable)
			}
		} else if eg.LeftExpressionGroup != nil {
			s += eg.LeftExpressionGroup.PrintAST(prefix, isLast)
		}
		return s
	}

	if isLast {
		s += fmt.Sprintf("%s└── %s\n", prefix, eg.Op)
	} else {
		s += fmt.Sprintf("%s├── %s\n", prefix, eg.Op)
	}

	newPrefix := prefix
	if isLast {
		newPrefix += "    "
	} else {
		newPrefix += "│   "
	}

	childrenCount := 0
	if eg.LeftVariable != nil || eg.LeftExpressionGroup != nil {
		childrenCount++
	}
	if eg.RightVariable != nil || eg.RightExpressionGroup != nil {
		childrenCount++
	}

	printedChildren := 0

	if eg.LeftVariable != nil {
		printedChildren++
		isLastChild := (printedChildren == childrenCount)
		if isLastChild {
			s += fmt.Sprintf("%s└── %s\n", newPrefix, eg.LeftVariable)
		} else {
			s += fmt.Sprintf("%s├── %s\n", newPrefix, eg.LeftVariable)
		}
	} else if eg.LeftExpressionGroup != nil {
		printedChildren++
		isLastChild := (printedChildren == childrenCount)
		s += eg.LeftExpressionGroup.PrintAST(newPrefix, isLastChild)
	}

	if eg.RightVariable != nil {
		printedChildren++
		isLastChild := (printedChildren == childrenCount)
		if isLastChild {
			s += fmt.Sprintf("%s└── %s\n", newPrefix, eg.RightVariable)
		} else {
			s += fmt.Sprintf("%s├── %s\n", newPrefix, eg.RightVariable)
		}
	} else if eg.RightExpressionGroup != nil {
		printedChildren++
		isLastChild := (printedChildren == childrenCount)
		s += eg.RightExpressionGroup.PrintAST(newPrefix, isLastChild)
	}

	return s
}

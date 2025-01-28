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
		if variable.Not {
			newValue = newValue.NOT()
		}
		return factManager.SetFactValueByLetter(variable.Letter, newValue, false)
	} else {
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
		return sideDeduction(ep.LeftVariable, ep.LeftExpressionGroup, newLeftValue)
	} else if newRightValue != v.UNKNOWN {
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

func (eg *ExpressionGroup) PrintAST(prefix string, isLast bool) {
    if eg.Op == NOTHING {
        if eg.LeftVariable != nil {
            if isLast {
                fmt.Printf("%s└── %s\n", prefix, eg.LeftVariable)
            } else {
                fmt.Printf("%s├── %s\n", prefix, eg.LeftVariable)
            }
        } else if eg.LeftExpressionGroup != nil {
            eg.LeftExpressionGroup.PrintAST(prefix, isLast)
        }
        return
    }

    if isLast {
        fmt.Printf("%s└── %s\n", prefix, eg.Op)
    } else {
        fmt.Printf("%s├── %s\n", prefix, eg.Op)
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

    if eg.LeftVariable != nil {
        if childrenCount == 2 {
            fmt.Printf("%s├── %s\n", newPrefix, eg.LeftVariable)
        } else {
            fmt.Printf("%s└── %s\n", newPrefix, eg.LeftVariable)
        }
    } else if eg.LeftExpressionGroup != nil {
        eg.LeftExpressionGroup.PrintAST(newPrefix, childrenCount < 2)
    }

    if eg.RightVariable != nil {
        fmt.Printf("%s└── %s\n", newPrefix, eg.RightVariable)
    } else if eg.RightExpressionGroup != nil {
        eg.RightExpressionGroup.PrintAST(newPrefix, true)
    }
}

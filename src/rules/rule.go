package rules

import (
	"expert-system/src/factManager"
	"expert-system/src/v"
	"fmt"
	"sort"
)

var RuleDisplayLogs bool = false

type Side int

const (
	LEFT Side = iota
	RIGHT
)

type Rule struct {
	Op                   ConditionalOperator
	LeftExpressionGroup  *ExpressionGroup
	RightExpressionGroup *ExpressionGroup
	LeftVariable         *Variable
	RightVariable        *Variable
}

func (rule Rule) Solving() (v.Value, v.Value, *v.Error) {
	LogRule(fmt.Sprintf("%s solving", rule))
	expressionGroupTmp := ExpressionGroup{
		Op:                   NOTHING,
		LeftVariable:         rule.LeftVariable,
		RightVariable:        rule.RightVariable,
		LeftExpressionGroup:  rule.LeftExpressionGroup,
		RightExpressionGroup: rule.RightExpressionGroup,
	}
	// LogRule(fmt.Sprintf("ExpressionGroupTmp: %s", expressionGroupTmp))
	// LogRule(fmt.Sprintf("expressionGroupTmp.leftVariable: %s", expressionGroupTmp.LeftVariable))
	// LogRule(fmt.Sprintf("expressionGroupTmp.rightVariable: %s", expressionGroupTmp.RightVariable))
	// LogRule(fmt.Sprintf("expressionGroupTmp.leftExpressionGroup: %s", expressionGroupTmp.LeftExpressionGroup))
	// LogRule(fmt.Sprintf("expressionGroupTmp.rightExpressionGroup: %s", expressionGroupTmp.RightExpressionGroup))
	leftValue, err := solvingSide(expressionGroupTmp.LeftVariable, expressionGroupTmp.LeftExpressionGroup)
	if err != nil {
		return v.UNDETERMINED, v.UNDETERMINED, err
	}
	rightValue, err := solvingSide(expressionGroupTmp.RightVariable, expressionGroupTmp.RightExpressionGroup)
	if err != nil {
		return v.UNDETERMINED, v.UNDETERMINED, err
	}
	LogRule(fmt.Sprintf("%s solving, LeftValue: %s, RightValue: %s", rule, leftValue, rightValue))

	if leftValue == v.TRUE && rightValue == v.FALSE {
		return leftValue, rightValue, &v.Error{Type: v.CONTRADICTION, Message: fmt.Sprintf("%s %s %s, for rule %s", leftValue, rule.Op, rightValue, rule)}
	}
	return leftValue, rightValue, nil
}

// func (rule Rule) RuleDeduction(leftValue v.Value, rightValue v.Value) *v.Error {
// 	LogRule(fmt.Sprintf("%s deduction, LeftValue: %s, RightValue: %s", rule, leftValue, rightValue))
// 	if leftValue.Real() && rightValue == v.UNKNOWN {
// 		if rule.RightVariable != nil {
//             if rule.RightVariable.Not {
//                 leftValue = leftValue.NOT()
//             }
// 			return factManager.SetFactValueByLetter(rule.RightVariable.Letter, leftValue, false)
// 		} else {
// 			return rule.RightExpressionGroup.deduction(leftValue)
// 		}
// 	} else if leftValue == v.UNKNOWN && rightValue.Real() {
// 		if rule.LeftVariable != nil {
//             if rule.LeftVariable.Not {
//                 rightValue = rightValue.NOT()
//             }
// 			return factManager.SetFactValueByLetter(rule.LeftVariable.Letter, rightValue, false)
// 		} else {
// 			return rule.LeftExpressionGroup.deduction(rightValue)
// 		}
// 	}
// 	return nil
// }

func (rule Rule) RuleDeduction(leftValue v.Value, rightValue v.Value) *v.Error {
	LogRule(fmt.Sprintf("%s deduction, LeftValue: %s, RightValue: %s", rule, leftValue, rightValue))
	if leftValue == v.TRUE && rightValue == v.UNKNOWN {
		if rule.RightVariable != nil {
            if rule.RightVariable.Not {
                leftValue = leftValue.NOT()
            }
			return factManager.SetFactValueByLetter(rule.RightVariable.Letter, leftValue, false)
		} else {
			return rule.RightExpressionGroup.deduction(leftValue)
		}
	} else if leftValue == v.TRUE && rightValue == v.FALSE {
		return &v.Error{Type: v.CONTRADICTION, Message: fmt.Sprintf("%s %s %s, for rule %s", leftValue, rule.Op, rightValue, rule)}
	}
	return nil
}

func RulesConditionalOperatorFormatter(rules []Rule) []Rule {
	var newRules []Rule
	for _, rule := range rules {
		if rule.Op == IFF {
			newRules = append(newRules, Rule{
				LeftExpressionGroup:  rule.LeftExpressionGroup,
				RightExpressionGroup: rule.RightExpressionGroup,
				LeftVariable:         rule.LeftVariable,
				RightVariable:        rule.RightVariable,
				Op:                   IMPLIES,
			})
			newRules = append(newRules, Rule{
				LeftExpressionGroup:  rule.RightExpressionGroup,
				RightExpressionGroup: rule.LeftExpressionGroup,
				LeftVariable:         rule.RightVariable,
				RightVariable:        rule.LeftVariable,
				Op:                   IMPLIES,
			})
		} else {
			newRules = append(newRules, rule)
		}
	}
	return newRules
}

// SortFactList - Sort factList by fact occurence in ruleList. Facts with UNKNOWN value are prioritized.
func SortFactList(ruleList []Rule, factList []factManager.Fact, lap int) []factManager.Fact {
	var factListOccurence = make(map[rune]int)
	for _, fact := range factList {
		factListOccurence[fact.Letter] = 0
	}
	for _, rule := range ruleList {
		if rule.LeftVariable != nil {
			factListOccurence[rule.LeftVariable.Letter]++
		} else {
			rule.LeftExpressionGroup.getFactOccurences(&factListOccurence)
		}
		if rule.RightVariable != nil {
			factListOccurence[rule.RightVariable.Letter]++
		} else {
			rule.RightExpressionGroup.getFactOccurences(&factListOccurence)
		}
	}
	sort.Slice(factList, func(i, j int) bool {
		occurrenceI := factListOccurence[factList[i].Letter]
		occurrenceJ := factListOccurence[factList[j].Letter]
		if (factList[i].Value == v.UNKNOWN && factList[j].Value != v.UNKNOWN) {
			return true
		} else if (factList[i].Value != v.UNKNOWN && factList[j].Value == v.UNKNOWN) {
			return false
		}
		if occurrenceI == occurrenceJ {
			return factList[i].Letter < factList[j].Letter
		}
		return occurrenceI > occurrenceJ
	})
	newFactList := make([]factManager.Fact, len(factList))
	copy(newFactList, factList)
	// ftm.Printf("Sort with lap %d\n", lap)
	for lap > 0 {
		newFactListTmp := []factManager.Fact{}
		newFactListTmp = append(newFactListTmp, newFactList[1:]...)
		newFactListTmp = append(newFactListTmp, newFactList[0])
		newFactList = newFactListTmp
		lap--
	}
	return newFactList
}

// DISPLAY

func LogRule(msg string) {
	if RuleDisplayLogs {
		fmt.Println(fmt.Sprintf("Rule - %s", msg))
	}
}

func (rule Rule) String() string {
	return fmt.Sprintf("%s %s %s", rule.DisplaySide(LEFT), rule.Op, rule.DisplaySide(RIGHT))
}

func (rule Rule) DisplaySide(side Side) string {
	if side == LEFT {
		if rule.LeftVariable != nil {
			return rule.LeftVariable.String()
		} else {
			return rule.LeftExpressionGroup.String()
		}
	} else {
		if rule.RightVariable != nil {
			return rule.RightVariable.String()
		} else {
			return rule.RightExpressionGroup.String()
		}
	}
}

func DisplayRules(rules []Rule) {
	fmt.Println("---------- RULES ----------")
	for i, rule := range rules {
		fmt.Printf("%d: %s\n", i, rule.String())
	}
}

func (r *Rule) PrintAST() {
    fmt.Printf("%s\n", r.Op)

    childrenCount := 0
    if r.LeftVariable != nil || r.LeftExpressionGroup != nil {
        childrenCount++
    }
    if r.RightVariable != nil || r.RightExpressionGroup != nil {
        childrenCount++
    }

    printedChildren := 0

    if r.LeftVariable != nil {
        printedChildren++
        isLastChild := (printedChildren == childrenCount)
        if isLastChild {
            fmt.Printf("└── %s\n", r.LeftVariable)
        } else {
            fmt.Printf("├── %s\n", r.LeftVariable)
        }
    } else if r.LeftExpressionGroup != nil {
        printedChildren++
        isLastChild := (printedChildren == childrenCount)
        r.LeftExpressionGroup.PrintAST("", isLastChild)
    }

    if r.RightVariable != nil {
        printedChildren++
        isLastChild := (printedChildren == childrenCount)
        if isLastChild {
            fmt.Printf("└── %s\n", r.RightVariable)
        } else {
            fmt.Printf("├── %s\n", r.RightVariable)
        }
    } else if r.RightExpressionGroup != nil {
        printedChildren++
        isLastChild := (printedChildren == childrenCount)
        r.RightExpressionGroup.PrintAST("", isLastChild)
    }
}

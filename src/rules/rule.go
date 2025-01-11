package rules

import (
	"errors"
	"expert/factManager"
	"expert/v"
	"fmt"
)

var RuleDisplayLogs bool = false

type Side int

const (
    LEFT Side = iota
    RIGHT
)

type Rule struct {
	Op ConditionalOperator
	LeftExpressionGroup *ExpressionGroup
	RightExpressionGroup *ExpressionGroup
	LeftVariable *Variable
	RightVariable *Variable
}

func (rule Rule) Solving() (v.Value, v.Value, error) {

	expressionGroupTmp := ExpressionGroup{
		Op: NOTHING,
		LeftVariable: rule.LeftVariable,
		RightVariable: rule.RightVariable,
		LeftExpressionGroup: rule.LeftExpressionGroup,
		RightExpressionGroup: rule.RightExpressionGroup,
	}

	leftValue, err := solvingSide(expressionGroupTmp.LeftVariable, expressionGroupTmp.LeftExpressionGroup)
	if err != nil {
		return v.UNDETERMINED, v.UNDETERMINED, err
	}
	rightValue, err := solvingSide(expressionGroupTmp.RightVariable, expressionGroupTmp.RightExpressionGroup)
	if err != nil {
		return v.UNDETERMINED, v.UNDETERMINED, err
	}
	LogRule(fmt.Sprintf("%s solving, LeftValue: %s, RightValue: %s", rule, leftValue, rightValue))

	if leftValue.Real() && rightValue.Real() && (leftValue != rightValue) {
		return leftValue, rightValue, errors.New(fmt.Sprintf("CONTRADICTION : %s %s %s, for rule %s", leftValue, rule.Op, rightValue, rule))
	}
	return leftValue, rightValue, nil
}

func (rule Rule) RuleDeduction(leftValue v.Value, rightValue v.Value) error {
	LogRule(fmt.Sprintf("%s deduction, LeftValue: %s, RightValue: %s", rule, leftValue, rightValue))
	if leftValue.Real() && rightValue == v.UNKNOWN {
		if rule.RightVariable != nil {
            if rule.RightVariable.Not {
                leftValue = leftValue.NOT()
            }
			return factManager.SetFactValueByLetter(rule.RightVariable.Letter, leftValue)
		} else {
			return rule.RightExpressionGroup.deduction(leftValue)
		}
	} else if leftValue == v.UNKNOWN && rightValue.Real() {
		if rule.LeftVariable != nil {
            if rule.LeftVariable.Not {
                rightValue = rightValue.NOT()
            }
			return factManager.SetFactValueByLetter(rule.LeftVariable.Letter, rightValue)
		} else {
			return rule.LeftExpressionGroup.deduction(rightValue)
		}
	}
	return nil
}

func RulesConditionalOperatorFormatter(rules []Rule) []Rule {
    var newRules []Rule
    for _, rule := range rules {
        if rule.Op == IFF {
            newRules = append(newRules, Rule{
                LeftExpressionGroup: rule.LeftExpressionGroup,
				RightExpressionGroup: rule.RightExpressionGroup,
				LeftVariable: rule.LeftVariable,
				RightVariable: rule.RightVariable,
				Op: IMPLIES,
            })
            newRules = append(newRules, Rule{
                LeftExpressionGroup: rule.RightExpressionGroup,
				RightExpressionGroup: rule.LeftExpressionGroup,
				LeftVariable: rule.RightVariable,
				RightVariable: rule.LeftVariable,
				Op: IMPLIES,
            })
        } else {
            newRules = append(newRules, rule)
        }
    }
    return newRules
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

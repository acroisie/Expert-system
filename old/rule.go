package rules

import (
	"errors"
	"fmt"
)

type Rule struct {
	Op ConditionalOperator
	LeftExpressionGroup *ExpressionGroup
	RightExpressionGroup *ExpressionGroup
	LeftVariable *Variable
	RightVariable *Variable
}

func (rule Rule) Solving(facts []Fact, display bool) (Value, Value, error) {
	if display {
		// fmt.Println("\nRule solving : ", rule)
	}

	expressionGroupTmp := ExpressionGroup{
		Op: NOTHING,
		LeftVariable: rule.LeftVariable,
		RightVariable: rule.RightVariable,
		LeftExpressionGroup: rule.LeftExpressionGroup,
		RightExpressionGroup: rule.RightExpressionGroup,
	}
	
	leftValue, err := expressionGroupTmp.solvingSide(true, facts)
	if err != nil {
		return v.UNDETERMINED, v.UNDETERMINED, err
	}

	rightValue, err := expressionGroupTmp.solvingSide(false, facts)
	if err != nil {
		return v.UNDETERMINED, v.UNDETERMINED, err
	}
	if display {
		fmt.Println(fmt.Sprintf("Rule %s solving, LeftValue: %s, RightValue: %s", rule, leftValue, rightValue))
	}

	// leftResult, RightResult, err := rule.Op.solve(leftValue, rightValue)
	// if err != nil {
	// 	return v.UNDETERMINED, v.UNDETERMINED, err
	// }

	// if display {
	// 	fmt.Println(fmt.Sprintf("Rule solving LeftResult: %s, RightResult: %s", leftResult, RightResult))
	// }
	if leftValue.Real() && rightValue.Real() && (leftValue != rightValue) {
		return leftValue, rightValue, errors.New(fmt.Sprintf("CONTRADICTION : %s %s %s, for rule %s", leftValue, rule.Op, rightValue, rule))
	}
	return leftValue, rightValue, nil
}

func SolveRules(rules []Rule, facts []Fact) {
	fmt.Println("\n\n---------- SOLVING RULES ----------")
	var change int = 1
	var lap int = 0
	for change > 0 && lap < 30 {
		fmt.Println(fmt.Sprintf("\n---------- LAP %d ----------", lap))
		change = 0
		for _, rule := range rules {
			leftResult, RightResult, err := rule.Solving(facts, true)
			if err != nil {
				fmt.Printf("Error: %s\n", err)
				return	;
			}
			// fmt.Println(fmt.Sprintf("Rule %s : LeftResult: %s, RightResult: %s", rule, leftResult, RightResult))
			changeTmp, err := rule.ExploitResults(leftResult, RightResult, facts, true)
			if err != nil {
				fmt.Printf("Error: %s\n", err)
				return	;
			}
			change += changeTmp
		}
		lap++
	}
}

func (rule Rule) ExploitResults(leftValue Value, rightValue Value, facts []Fact, display bool) (int, error) {
	if leftValue.Real() && rightValue == v.UNKNOWN {
		if rule.RightVariable != nil {
			err := SetFactValueByLetter(facts, *rule.RightVariable, leftValue, true)
			return 1, err
		} else {
			return rule.RightExpressionGroup.resultDeduction(leftValue, 0, facts)
		}
	} else if leftValue == v.UNKNOWN && rightValue.Real() {
		if rule.LeftVariable != nil {
			err := SetFactValueByLetter(facts, *rule.LeftVariable, rightValue, true)
			return 1, err
		} else {
			return rule.LeftExpressionGroup.resultDeduction(rightValue, 0, facts)
		}
	}
	return 0, nil
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

func (rule Rule) String() string {
	return fmt.Sprintf("%s %s %s", rule.SideDisplay(true), rule.Op, rule.SideDisplay(false))
}

func (rule Rule) SideDisplay(side bool) string {
	if side {
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

// MOCK

func GetRulesMock() []Rule {
    rules := []Rule{
        // C => E
        Rule{
            Op: IMPLIES,
            LeftExpressionGroup: nil,
            LeftVariable: &Variable{Letter: 'C', Not: false},
            RightExpressionGroup: nil,
            RightVariable: &Variable{Letter: 'E', Not: false},
        },
        
        // A + B + C => D
        Rule{
            Op: IMPLIES,
            LeftExpressionGroup: &ExpressionGroup{
                Op: AND,
                LeftVariable: &Variable{Letter: 'A', Not: false},
                RightExpressionGroup: &ExpressionGroup{
                    Op: AND,
                    LeftVariable: &Variable{Letter: 'B', Not: false},
                    RightVariable: &Variable{Letter: 'C', Not: false},
                },
            },
            RightExpressionGroup: nil,
            RightVariable: &Variable{Letter: 'D', Not: false},
        },

        // A | B => C
        Rule{
            Op: IMPLIES,
            LeftExpressionGroup: &ExpressionGroup{
                Op: OR,
                LeftVariable: &Variable{Letter: 'A', Not: false},
                RightVariable: &Variable{Letter: 'B', Not: false},
            },
            RightExpressionGroup: nil,
            RightVariable: &Variable{Letter: 'C', Not: false},
        },

        // A + !B => F
        Rule{
            Op: IMPLIES,
            LeftExpressionGroup: &ExpressionGroup{
                Op: AND,
                LeftVariable: &Variable{Letter: 'A', Not: false},
                RightVariable: &Variable{Letter: 'B', Not: true}, // NOT appliqué
            },
            RightExpressionGroup: nil,
            RightVariable: &Variable{Letter: 'F', Not: false},
        },

        // C | !G => H
        Rule{
            Op: IMPLIES,
            LeftExpressionGroup: &ExpressionGroup{
                Op: OR,
                LeftVariable: &Variable{Letter: 'C', Not: false},
                RightVariable: &Variable{Letter: 'G', Not: true},
            },
            RightExpressionGroup: nil,
            RightVariable: &Variable{Letter: 'H', Not: false},
        },

        // V ^ W => X
        Rule{
            Op: IMPLIES,
            LeftExpressionGroup: &ExpressionGroup{
                Op: XOR,
                LeftVariable: &Variable{Letter: 'V', Not: false},
                RightVariable: &Variable{Letter: 'W', Not: false},
            },
            RightExpressionGroup: nil,
            RightVariable: &Variable{Letter: 'X', Not: false},
        },

        // A + B => Y + Z
        Rule{
            Op: IMPLIES,
            LeftExpressionGroup: &ExpressionGroup{
                Op: AND,
                LeftVariable: &Variable{Letter: 'A', Not: false},
                RightVariable: &Variable{Letter: 'B', Not: false},
            },
            RightExpressionGroup: &ExpressionGroup{
                Op: AND,
                LeftVariable: &Variable{Letter: 'Y', Not: false},
                RightVariable: &Variable{Letter: 'Z', Not: false},
            },
        },

        // C | D => X | V
        Rule{
            Op: IMPLIES,
            LeftExpressionGroup: &ExpressionGroup{
                Op: OR,
                LeftVariable: &Variable{Letter: 'C', Not: false},
                RightVariable: &Variable{Letter: 'D', Not: false},
            },
            RightExpressionGroup: &ExpressionGroup{
                Op: OR,
                LeftVariable: &Variable{Letter: 'X', Not: false},
                RightVariable: &Variable{Letter: 'V', Not: false},
            },
        },

        // E + F => !V
        Rule{
            Op: IMPLIES,
            LeftExpressionGroup: &ExpressionGroup{
                Op: AND,
                LeftVariable: &Variable{Letter: 'E', Not: false},
                RightVariable: &Variable{Letter: 'F', Not: false},
            },
            RightExpressionGroup: nil,
            RightVariable: &Variable{Letter: 'V', Not: true},
        },

        // A + B <=> C
        Rule{
            Op: IFF,
            LeftExpressionGroup: &ExpressionGroup{
                Op: AND,
                LeftVariable: &Variable{Letter: 'A', Not: false},
                RightVariable: &Variable{Letter: 'B', Not: false},
            },
            RightExpressionGroup: nil,
            RightVariable: &Variable{Letter: 'C', Not: false},
        },

        // A + B <=> !C
        // Rule{
        //     Op: IFF,
        //     LeftExpressionGroup: &ExpressionGroup{
        //         Op: AND,
        //         LeftVariable: &Variable{Letter: 'A', Not: false},
        //         RightVariable: &Variable{Letter: 'B', Not: false},
        //     },
        //     RightExpressionGroup: nil,
        //     RightVariable: &Variable{Letter: 'C', Not: true},
        // },
    }
    return rules
}



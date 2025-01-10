package main

import (
    "fmt"
    "expert/rules"
    "expert/factManager"
)

func algo(ruleList []rules.Rule) {
	factManager.FactDisplayLogs = true
	rules.RuleDisplayLogs = true
	rules.ExpressionGroupDisplayLogs = true
	
    fmt.Println("\n\n---------- SOLVING RULES ----------")
	var lap int = 0
    factManager.FactChangeCounter = 1

	for factManager.FactChangeCounter > 0 && lap < 30 {

		fmt.Println(fmt.Sprintf("\n---------- LAP %d ----------", lap))
		factManager.FactChangeCounter = 0

		for _, rule := range ruleList {

			leftResult, RightResult, err := rule.Solving()
			if err != nil {
				fmt.Printf("Error: %s\n", err)
				return	;
			}

			err = rule.RuleDeduction(leftResult, RightResult)
			if err != nil {
				fmt.Printf("Error: %s\n", err)
				return	;
			}
		}
		lap++
	}
    fmt.Println(fmt.Sprintf("\n---------- END OF ALGO - %d LAP ----------", lap))
    factManager.Display()
}

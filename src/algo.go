package main

import (
    "fmt"
    "expert-system/src/rules"
    "expert-system/src/factManager"
)

var RuleList []rules.Rule

func Algo(ruleList []rules.Rule) {
	fmt.Println("\n\n---------- ALGO ----------")
	factManager.FactDisplayLogs = false
	rules.RuleDisplayLogs = false
	rules.ExpressionGroupDisplayLogs = false

	rules.DisplayRules(ruleList)
	fmt.Println("\n\n---------- INITIAL FACTS ----------")
	factManager.Display()

	RuleList = ruleList
	ForwardChecking()
	fmt.Println("\n\n---------- FACTS AFTER FORWARD CHECKING ----------")
	factManager.Display()

	unknowLetters := factManager.GetUnknowLetters()
	if len(unknowLetters) > 0 {
		result, err := BackTracking()
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(result)
		}
	}
}

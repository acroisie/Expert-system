package algo

import (
    "fmt"
    "expert-system/src/rules"
    "expert-system/src/factManager"
)

var RuleList []rules.Rule
var AlgoDisplayLogs bool = false

func Algo(ruleList []rules.Rule) {
	AlgoDisplayLogs = false
	factManager.FactDisplayLogs = false
	rules.RuleDisplayLogs = false
	rules.ExpressionGroupDisplayLogs = false

	if AlgoDisplayLogs {
		fmt.Println("\n\n---------- ALGO ----------")
		rules.DisplayRules(ruleList)
		fmt.Println("\n\n---------- INITIAL FACTS ----------")
		factManager.Display()
	}

	RuleList = ruleList
	FCError := forwardChecking()
	if FCError != nil {
		fmt.Printf(FCError.Message)
		return
	}
	if AlgoDisplayLogs {
		fmt.Println("\n\n---------- FACTS AFTER FORWARD CHECKING ----------")
		factManager.Display()
	}

	unknowLetters := factManager.GetUnknowLetters()
	if len(unknowLetters) > 0 {
		BTError := backTracking()
		if BTError != nil {
			fmt.Printf(BTError.Message)
			return
		}
	}

	fmt.Printf("\n\nFACT LIST RESULT:\n")
	factManager.Display()
}

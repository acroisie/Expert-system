package algo

import (
    "fmt"
    "expert-system/src/rules"
    "expert-system/src/factManager"
	"expert-system/src/v"
)

var RuleList []rules.Rule
var AlgoDisplayLogs bool = false

func Algo(ruleList []rules.Rule) bool {
	AlgoDisplayLogs = true
	factManager.FactDisplayLogs = true
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
		return false
	}
	if AlgoDisplayLogs {
		fmt.Println("\n\n---------- FACTS AFTER FORWARD CHECKING ----------")
		factManager.Display()
	}

	onlyLeftError := SetFalseOnlyLeftFact()
	if onlyLeftError == nil {
		return true
	}

	unknowLetters := factManager.GetUnknowLetters()
	if len(unknowLetters) > 0 {
		BTError := backTracking()
		if BTError != nil {
			fmt.Printf(BTError.Message)
			return false
		}
	}

	return true
}

func SetFalseOnlyLeftFact() *v.Error {
	onlyLeftLetters := rules.SetLeftOnlyFacts(RuleList)
	if AlgoDisplayLogs {
		fmt.Println("\n\n---------- ONLY LEFT FACTS ----------")
		stringTmp := ""
		for _, letter := range onlyLeftLetters {
			stringTmp += fmt.Sprintf("%c ", letter)
		}
		fmt.Println(stringTmp)
	}
	for _, letter := range onlyLeftLetters {
		err := factManager.SetFactValueByLetter(letter, v.FALSE, false)
		if err != nil && AlgoDisplayLogs {
			fmt.Printf("%s\n", err.Message)
		}
	}
	ruleListSave := make([]rules.Rule, len(RuleList))
	copy(ruleListSave, RuleList)

	FCError := forwardChecking()
	if FCError != nil {
		if AlgoDisplayLogs {
			fmt.Printf("\n\n---------- INIT ONLY LEFT FACTS TO FALSE - ERROR: %s ----------\n", FCError.Message)
		}
		RuleList = ruleListSave
		return FCError
	}
	unknownLetters := factManager.GetUnknowLetters()
	RuleList = ruleListSave
	if len(unknownLetters) > 0 {
		if AlgoDisplayLogs {
			fmt.Println("\n\n---------- SOLVED AFTER INIT ONLY LEFT FACTS TO FALSE - FAILURE ----------")
		}
		return &v.Error{Type: v.FACT_NOT_FOUND, Message: "Can't solve after init only left facts to false"}
	}
	if AlgoDisplayLogs {
		fmt.Println("\n\n---------- SOLVED AFTER INIT ONLY LEFT FACTS TO FALSE - SUCCESS ----------")
	}
	return nil
}
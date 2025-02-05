package algo

import (
	"expert-system/src/factManager"
	"expert-system/src/rules"
	"expert-system/src/v"
	"fmt"
)

var RuleList []rules.Rule
var AlgoDisplayLogs bool = false

func Algo(ruleList []rules.Rule) (bool, []string) {
	AlgoDisplayLogs = false
	factManager.FactDisplayLogs = false
	rules.RuleDisplayLogs = false
	rules.ReasoningDisplayLogs = true
	rules.ExpressionGroupDisplayLogs = false
	rules.ReasoningLogs = []string{}

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
		return false, []string{}
	}
	if AlgoDisplayLogs {
		fmt.Println("\n\n---------- FACTS AFTER FORWARD CHECKING ----------")
		factManager.Display()
	}
	logs := rules.ReasoningLogs
	rules.ReasoningLogs = []string{}

	setFalseByDefaultError := SetFalseByDefault()
	for _, log := range rules.ReasoningLogs {
		logs = append(logs, log)
	}
	if setFalseByDefaultError == nil {
		return true, logs
	}
	rules.ReasoningLogs = []string{}

	unknowLetters := factManager.GetUnknowLetters()
	if len(unknowLetters) > 0 {
		BTError := backTracking()
		if BTError != nil {
			fmt.Printf(BTError.Message)
			return false, []string{}
		}
	}

	falseFacts := ""
	trueFacts := ""
	for _, unknownLetter := range unknowLetters {
		fact, err := factManager.GetFactReferenceByLetter(unknownLetter)
		if err != nil {
			fmt.Printf(err.Message)
			return false, []string{}
		}
		if fact.Value == v.TRUE {
			trueFacts += fmt.Sprintf("%c ", unknownLetter)
		} else if fact.Value == v.FALSE {
			falseFacts += fmt.Sprintf("%c ", unknownLetter)
		}
	}
	if falseFacts != "" {
		logs = append(logs, fmt.Sprintf("%s= FALSE by default\n", falseFacts))
	}
	if trueFacts != "" {
		logs = append(logs, fmt.Sprintf("%s= TRUE by deduction\n", trueFacts))
	}

	return true, logs
}

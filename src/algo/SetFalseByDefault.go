package algo

import (
    "fmt"
    "expert-system/src/rules"
    "expert-system/src/factManager"
	"expert-system/src/v"
)

func SetFalseByDefault() *v.Error {
	onlyLeftErr := SetFalseOnlyLeftFact()
	if onlyLeftErr == nil {
		return nil
	} else if onlyLeftErr.Type != v.UNKNOWN_PRESENCE {
		return onlyLeftErr
	}

	inactiveLetters := true
	for inactiveLetters {
		inactiveLettersPresence, err := SetFalseInactiveFacts()
		if err != nil {
			return err
		}
		if !inactiveLettersPresence {
			inactiveLetters = false
			break
		}
	}

	unknownLetters := factManager.GetUnknowLetters()
	if len(unknownLetters) > 0 {
		return &v.Error{Type: v.UNKNOWN_PRESENCE, Message: "Can't solve after init inactive facts to false"}
	}
	return nil
}

func SetFalseInactiveFacts() (bool, *v.Error) {
	unknowLetters := factManager.GetUnknowLetters()
	activeLetters := make(map[rune]bool)

	for _, rule := range RuleList {
		leftResult, _, err := rule.Solving()
		if err != nil {
			return false, err
		}
		if leftResult != v.FALSE {
			if rule.RightVariable != nil {
				activeLetters[rule.RightVariable.Letter] = true
			} else if rule.RightExpressionGroup != nil {
				rightLetters:= rule.RightExpressionGroup.GetLetters()
				for letter := range rightLetters {
					activeLetters[letter] = true
				}
			}
		}
	}

	inactiveLetters := []rune{}
	for _, letter := range unknowLetters {
		if !activeLetters[letter] {
			inactiveLetters = append(inactiveLetters, letter)
		}
	}
	for _, letter := range inactiveLetters {
		err := factManager.SetFactValueByLetter(letter, v.FALSE, false)
		if err == nil {
			rules.LogReasoning(fmt.Sprintf("%c is present on the premises parts and on the conclusion parts, but the rules are not active, so %c = FALSE by default\n", letter, letter))
		}
	}

	FCError := forwardChecking()
	if FCError != nil {
		return false, FCError
	}

	if len(inactiveLetters) > 0 {
		return true, nil
	}
	return false, nil
}

func SetFalseOnlyLeftFact() *v.Error {
	onlyLeftLetters := rules.GetLeftOnlyFacts(RuleList)

	for _, letter := range onlyLeftLetters {
		err := factManager.SetFactValueByLetter(letter, v.FALSE, false)
		if err == nil {
			rules.LogReasoning(fmt.Sprintf("%c is only present on the premises parts, so %c = FALSE by default\n", letter, letter))
		}
	}
	ruleListSave := make([]rules.Rule, len(RuleList))
	copy(ruleListSave, RuleList)

	displayLogSave := rules.ReasoningDisplayLogs
	rules.ReasoningDisplayLogs = false
	FCError := forwardChecking()
	if FCError != nil {
		RuleList = ruleListSave
		return FCError
	}
	rules.ReasoningDisplayLogs = displayLogSave

	unknownLetters := factManager.GetUnknowLetters()
	RuleList = ruleListSave
	if len(unknownLetters) > 0 {
		return &v.Error{Type: v.UNKNOWN_PRESENCE, Message: "Can't solve after init only left facts to false"}
	}
	return nil
}
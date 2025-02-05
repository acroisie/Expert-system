package algo

import (
	"expert-system/src/factManager"
	"expert-system/src/v"
	"fmt"
)

func forwardChecking() *v.Error {

	var lap int = 0
	factManager.FactChangeCounter = 1

	for factManager.FactChangeCounter > 0 && lap < 30 {

		factManager.FactChangeCounter = 0

		for _, rule := range RuleList {

			leftResult, RightResult, err := rule.Solving()
			if err != nil {
				return &v.Error{Type: err.Type, Message: fmt.Sprintf("Error in the Forward Checking: %s", err.Message)}
			}

			err = rule.RuleDeduction(leftResult, RightResult)
			if err != nil {
				return &v.Error{Type: err.Type, Message: fmt.Sprintf("Error in the Forward Checking: %s", err.Message)}
			}
		}
		lap++
	}
	return nil
}

package main

import (
    "fmt"
    "expert-system/src/factManager"
	"expert-system/src/v"
)

func ForwardChecking() *v.Error {
	
    // fmt.Println("\n\n---------- FORWARD CHECKING ----------")
	var lap int = 0
    factManager.FactChangeCounter = 1

	for factManager.FactChangeCounter > 0 && lap < 30 {

		// fmt.Println(fmt.Sprintf("\n---------- LAP %d ----------", lap))
		factManager.FactChangeCounter = 0

		for _, rule := range RuleList {

			leftResult, RightResult, err := rule.Solving()
			if err != nil {
				fmt.Printf("Error: %s\n", err)
				return err
			}

			err = rule.RuleDeduction(leftResult, RightResult)
			if err != nil {
				fmt.Printf("Error: %s\n", err)
				return	err
			}
		}
		lap++
	}
    // fmt.Println(fmt.Sprintf("\n---------- END OF ALGO - %d LAP ----------", lap))
    // factManager.Display()
	return nil
}

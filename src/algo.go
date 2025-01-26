package main

import (
    "fmt"
    "expert-system/src/rules"
    "expert-system/src/factManager"
	"expert-system/src/v"
)

func Algo(ruleList []rules.Rule) {
	fmt.Println("\n\n---------- ALGO ----------")
	fmt.Println("\n\n---------- INITIAL FACTS ----------")
	factManager.Display()
	ForwardChecking(ruleList)
	fmt.Println("\n\n---------- FACTS AFTER FORWARD CHECKING ----------")
	factManager.Display()

	fmt.Print("\n\n---------- UNKNOW TO FALSE -> FORWARD CHECKING ----------")
	factListSave := make([]factManager.Fact, len(factManager.FactList))
	copy(factListSave, factManager.FactList)
	factManager.SetUnknowLettersToFalse()
	FCError := ForwardChecking(ruleList)
	fmt.Print("ForwardChecking result: ", FCError)
	fmt.Printf("With: \n")
	factManager.Display()
	if FCError == nil {
		fmt.Print("\n\n---------- SOLVED ----------")
		factManager.Display()
		return
	} else {
		fmt.Print("\n\n---------- BACKTRACKING ----------")
		factManager.FactList = factListSave
		fmt.Printf("factListSave: %v\n", factListSave)
		BackTracking(ruleList)
	}
}

func ForwardChecking(ruleList []rules.Rule) *v.Error {
	
    // fmt.Println("\n\n---------- FORWARD CHECKING ----------")
	var lap int = 0
    factManager.FactChangeCounter = 1

	for factManager.FactChangeCounter > 0 && lap < 30 {

		// fmt.Println(fmt.Sprintf("\n---------- LAP %d ----------", lap))
		factManager.FactChangeCounter = 0

		for _, rule := range ruleList {

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

func BackTracking(ruleList []rules.Rule) {
	factManager.FactDisplayLogs = false
	rules.RuleDisplayLogs = false
	rules.ExpressionGroupDisplayLogs = false
	
	factListSave := factManager.FactList
	unknowLetters := factManager.GetUnknowLetters()

	fmt.Println("\n\n---------- BACKTRACKING ----------")
	lap := 0
	for lap < len(unknowLetters) {
		fmt.Printf("\n---------- LAP %d ----------\n", lap)
		factListSorted := rules.SortFactList(ruleList, factListSave, lap)
		factManager.FactList = factListSorted

		unknowLettersChecker := make(map[rune]int)
		for _, letter := range unknowLetters {
			unknowLettersChecker[letter] = 0
		}

		fmt.Print("Unknow letters: ")
		factManager.DisplayRunesTab(unknowLetters)
		fmt.Println("FactList:")
		factManager.Display()
		fmt.Printf("len(unknowLetters) = %d\n", len(unknowLetters))

		endLoop := false
		for i := 0; i < len(unknowLetters); {
			fmt.Printf("i = %d\n", i)
			letter := unknowLetters[i]
			if unknowLettersChecker[letter] == 0 {
				factManager.SetFactValueByLetter(letter, v.FALSE, true)
			} else if unknowLettersChecker[letter] == 1 {
				factManager.SetFactValueByLetter(letter, v.TRUE, true)
			} else {
				fmt.Printf("Error: %c already checked\n", letter)
				if (i <= 0) {
					endLoop = true
					break
				} else {
					i--
				}
			}

			letterFact, err := factManager.GetFactReferenceByLetter(letter)
			if err != nil {
				fmt.Printf("Error: %s\n", err)
				return 
			}
			fmt.Println(fmt.Sprintf("Letter %c set to %s", letter, letterFact.Value))
			unknowLettersChecker[letter]++
			FCError := ForwardChecking(ruleList)
			fmt.Print(fmt.Sprintf("ForwardChecking result for letter %c = %s: %s\n", letter, letterFact.Value, FCError))
			factManager.Display()
			if FCError != nil {
				if (i <= 0) {
					i = 0
				} else {
					i--
				}
			} else {
				i++
			}
			if (i == 3) {
				return
			}
		}
		fmt.Println(fmt.Sprintf("End of lap with endLoop: %t", endLoop))
		return
		if endLoop == false {
			fmt.Println(fmt.Sprintf("LETS GO!"))
			break
		}
		lap++
	}
	factManager.Display()
}
package algo

import (
    "fmt"
    "expert-system/src/factManager"
	"expert-system/src/v"
)

func getFactPossibilities(unknowLetters []rune) ([][]factManager.Fact, *v.Error) {

	var factPossibilities = [][]factManager.Fact{}

	if AlgoDisplayLogs {
		fmt.Print("getFactPossibilities with Unknow letters: ")
		factManager.DisplayRunesTab(unknowLetters)
		fmt.Println("And FactList:")
		factManager.DisplayFactsOneLine(factManager.FactList)
	}

	unknowLettersChecker := make(map[rune]int)
	// 0: no value was checked; 1: FALSE value was checked; 2: TRUE value was checked
	for _, letter := range unknowLetters {
		unknowLettersChecker[letter] = 0
	}

	for i := 0; i < len(unknowLetters); {
		letter := unknowLetters[i]

		// if unknowLettersChecker[letter] > 1, it means that the letter has already been checked with TRUE and FALSE values
		if unknowLettersChecker[letter] > 1 {
			// reset the value of the letter to UNKNOWN
			unknowLettersChecker[letter] = 0
			factManager.SetFactValueByLetter(letter, v.UNKNOWN, true)
			// go back to the previous letter. If i == 0, it means that all possibilities have been checked
			if (i <= 0) {
				break
			} else {
				i--
			}
		} else {
			FCError := SetFactValueAndRunForwardChecking(unknowLettersChecker[letter], letter)
			unknowLettersChecker[letter]++
			// if FCERROR != nil, it means that the value of the letter is not compatible with the rules.
			// in this case, we re-check the letter with the other value
			if FCError == nil {
				// if FCError == nil, it means that the value of the letter is compatible with the rules.
				// we can go to the next letter, or if all letters have been checked, we add the factList to the possibilities
				// and go back to the previous letter to check if all his values have been checked
				if i < len(unknowLetters) - 1 {
					i++	
				} else {
					newPossibility := make([]factManager.Fact, len(factManager.FactList))
					copy(newPossibility, factManager.FactList)
					if AlgoDisplayLogs {
						fmt.Println("New possibility:")
						factManager.DisplayFactsOneLine(newPossibility)
					}

					factPossibilities = append(factPossibilities, newPossibility) 

					// unknowLettersChecker[letter] = 0
					// factManager.SetFactValueByLetter(letter, v.UNKNOWN, true)

					// i--
					break
				}
			}
		}
	}
	return factPossibilities, nil
}

func SetFactValueAndRunForwardChecking(checkerValue int, letter rune) *v.Error {
	if checkerValue == 0 {
		factManager.SetFactValueByLetter(letter, v.FALSE, true)
	} else if checkerValue == 1 {
		factManager.SetFactValueByLetter(letter, v.TRUE, true)
	}
	letterFact, err := factManager.GetFactReferenceByLetter(letter)
	if err != nil {
		return err
	}
	
	FCError := forwardChecking()
	if AlgoDisplayLogs {
		fmt.Print(fmt.Sprintf("ForwardChecking result for letter %c = %s: %s\n", letter, letterFact.Value, FCError))
	}
	return FCError
}






package main

import (
    "fmt"
    "expert-system/src/rules"
    "expert-system/src/factManager"
	"expert-system/src/v"
)

func factPossibilitiesFusion(factPossibilities [][]factManager.Fact) ([]factManager.Fact, *v.Error) {
    fmt.Println("FactPossibilitiesFusion")

    mergedFacts := make([]factManager.Fact, len(factManager.FactList))
    copy(mergedFacts, factManager.FactList)

    if len(factPossibilities) <= 0 {
        return mergedFacts, &v.Error{Type: v.SOLVING, Message: "No fact possibilities to merge"}
    }

    for i, fact := range mergedFacts {
        letter := fact.Letter

        allValues := make([]v.Value, 0, len(factPossibilities))
        for _, possibility := range factPossibilities {
            for _, factInPossibility := range possibility {
                if factInPossibility.Letter == letter {
                    allValues = append(allValues, factInPossibility.Value)
                    break
                }
            }
        }

        alwaysFalse := true
        alwaysTrue := true
        for _, val := range allValues {
            if val != v.FALSE {
                alwaysFalse = false
            }
            if val != v.TRUE {
                alwaysTrue = false
            }
        }

        if alwaysFalse {
            mergedFacts[i].Value = v.FALSE
        } else if alwaysTrue {
            mergedFacts[i].Value = v.TRUE
        } else {
			fmt.Printf("Letter %c is undetermined\n", letter)
			var possibilitiesWithFalseValue = [][]factManager.Fact{}
			var possibilitiesWithTrueValue = [][]factManager.Fact{}
            for _, possibility := range factPossibilities {
				value, err := factManager.GetLetterValue(possibility, letter)
				if err != nil {
					return mergedFacts, err
				}
				possibilityWithoutLetter := factManager.SliceALetter(possibility, letter)
				if value == v.FALSE {
					possibilitiesWithFalseValue = append(possibilitiesWithFalseValue, possibilityWithoutLetter)
				} else if value == v.TRUE {
					possibilitiesWithTrueValue = append(possibilitiesWithTrueValue, possibilityWithoutLetter)
				}
			}
			fmt.Printf("Possibilities with False value for letter %c: %d\n", letter, len(possibilitiesWithFalseValue))
			fmt.Printf("Possibilities with True value for letter %c: %d\n", letter, len(possibilitiesWithTrueValue))
			var possibilitiesWithTrueValueToKeep = [][]factManager.Fact{}
			for _, truePossibility := range possibilitiesWithTrueValue {
				fmt.Println("Possibility with True value:")
				factManager.DisplayFacts(truePossibility)
				toRemove := false
				for _, falsePossibility := range possibilitiesWithFalseValue {
					tmp := factManager.CompareFactLists(truePossibility, falsePossibility)
					if tmp {
						toRemove = true
						break
					}
				}
				if !toRemove {
					possibilitiesWithTrueValueToKeep = append(possibilitiesWithTrueValueToKeep, truePossibility)
				}
			}
			fmt.Printf("Possibilities with True value to keep for letter %c: %d\n", letter, len(possibilitiesWithTrueValueToKeep))
			if len(possibilitiesWithTrueValueToKeep) > 0 {
				mergedFacts[i].Value = v.UNDETERMINED
			} else {
				mergedFacts[i].Value = v.FALSE
			}
        }
    }

    return mergedFacts, nil
}

func factPossibilitiesFusion3(factPossibilities [][]factManager.Fact) ([]factManager.Fact, *v.Error) {
    fmt.Println("FactPossibilitiesFusion3")

    mergedFacts := make([]factManager.Fact, len(factManager.FactList))
    copy(mergedFacts, factManager.FactList)

    if len(factPossibilities) <= 0 {
        return mergedFacts, &v.Error{Type: v.SOLVING, Message: "No fact possibilities to merge"}
    }

    for i, fact := range mergedFacts {
        letter := fact.Letter

        allValues := make([]v.Value, 0, len(factPossibilities))
        for _, possibility := range factPossibilities {
            for _, factInPossibility := range possibility {
                if factInPossibility.Letter == letter {
                    allValues = append(allValues, factInPossibility.Value)
                    break
                }
            }
        }

        alwaysFalse := true
        alwaysTrue := true
        for _, val := range allValues {
            if val != v.FALSE {
                alwaysFalse = false
            }
            if val != v.TRUE {
                alwaysTrue = false
            }
        }

        if alwaysFalse {
            mergedFacts[i].Value = v.FALSE
        } else if alwaysTrue {
            mergedFacts[i].Value = v.TRUE
        } else {
			mergedFacts[i].Value = v.UNDETERMINED
        }
    }

    return mergedFacts, nil
}

func countTrueValues(factList []factManager.Fact) int {
	count := 0
	for _, fact := range factList {
		if fact.Value == v.TRUE {
			count++
		}
	}
	return count
}

func getSmallestPossibilities(factPossibilities [][]factManager.Fact) ([][]factManager.Fact) {
	smallestPossibilities := [][]factManager.Fact{}
	smallestPossibilitiesSize := 0
	for _, possibility := range factPossibilities {
		trueCounter := countTrueValues(possibility)
		if trueCounter < smallestPossibilitiesSize || smallestPossibilitiesSize == 0 {
			smallestPossibilities = [][]factManager.Fact{}
			smallestPossibilities = append(smallestPossibilities, possibility)
			smallestPossibilitiesSize = trueCounter
		} else if trueCounter == smallestPossibilitiesSize {
			smallestPossibilities = append(smallestPossibilities, possibility)
		}
	}
	return smallestPossibilities
}

func factPossibilitiesFusion2(factPossibilities [][]factManager.Fact) ([]factManager.Fact, *v.Error) {
	same := true
	for i := 0; i < len(factPossibilities); {
		for j := i + 1; j < len(factPossibilities); {
			if !factManager.CompareFactLists(factPossibilities[i], factPossibilities[j]) {
				same = false
				break
			}
			j++
		}
		if !same {
			i++
		} else {
			break
		}
	}
	if same {
		return factPossibilities[0], nil
	} else {
		return nil, &v.Error{Type: v.SOLVING, Message: "Multiple possibilities"}
	}
}

func SetFactValueAndLaunchForwardChecking(checkerValue int, letter rune) *v.Error {
	if checkerValue == 0 {
		factManager.SetFactValueByLetter(letter, v.FALSE, true)
	} else if checkerValue == 1 {
		factManager.SetFactValueByLetter(letter, v.TRUE, true)
	}
	letterFact, err := factManager.GetFactReferenceByLetter(letter)
	if err != nil {
		return err
	}
	
	FCError := ForwardChecking()
	fmt.Print(fmt.Sprintf("ForwardChecking result for letter %c = %s: %s\n", letter, letterFact.Value, FCError))
	return FCError
}

func getFactPossibilities(unknowLetters []rune) ([][]factManager.Fact, *v.Error) {

	var factPossibilities = [][]factManager.Fact{}

	fmt.Print("getFactPossibilities with Unknow letters: ")
	factManager.DisplayRunesTab(unknowLetters)
	fmt.Println("And FactList:")
	factManager.DisplayFactsOneLine(factManager.FactList)

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
			FCError := SetFactValueAndLaunchForwardChecking(unknowLettersChecker[letter], letter)
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
					fmt.Println("New possibility:")
					factManager.DisplayFactsOneLine(newPossibility)

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

func BackTracking() (*[]factManager.Fact, *v.Error) {
	fmt.Println("\n\n---------- BACKTRACKING ----------")
	
	factListSave := make([]factManager.Fact, len(factManager.FactList))
	factPossibilitiesTotal := [][]factManager.Fact{}
	copy(factListSave, factManager.FactList)
	lap := 0
	maxLap := len(factManager.GetUnknowLetters())
	for lap < maxLap {

		fmt.Printf("\n---------- LAP %d ----------\n", lap)
		// Sort factList by fact occurence in ruleList. Facts with UNKNOWN value are prioritized.
		factManager.FactList = rules.SortFactList(RuleList, factListSave, lap)
		// Get letters with UNKNOWN value, sorted by occurence in ruleList
		unknowLetters := factManager.GetUnknowLetters()

		factPossibilities, _ := getFactPossibilities(unknowLetters)
		for _, possibility := range factPossibilities {
			// fmt.Printf("Possibility %d:\n", e)
			// factManager.DisplayFacts(possibility)
			factPossibilitiesTotal = append(factPossibilitiesTotal, possibility)
		}
	
		// fmt.Printf("AnalyzeResult len = %d\n", len(factPossibilities))
		// factPossibilitiesFusionResult, err := factPossibilitiesFusion(factPossibilities)
		// fmt.Printf("FactPossibilitiesFusionResult len = %d\n", len(factPossibilitiesFusionResult))
		// return nil, err
		// if err != nil {
		// 	return nil, err
		// }
		// if len(analyzeResult) > 0 {
		// 	for _, result := range analyzeResult {
        //         factPossibilities = append(factPossibilities, result)
        //     }
		// }
		lap++
	}

	fmt.Println("\n\n---------- FACTS AFTER BACKTRACKING ----------")
	fmt.Printf("factPossibilitiesTotal len = %d\n", len(factPossibilitiesTotal))
	for i, factList := range factPossibilitiesTotal {
		fmt.Printf("\n---------- FACTS %d ----------\n", i)
		factManager.DisplayFacts(factList)
	}
	// factPossibilitiesTotalTmp := getSmallestPossibilities(factPossibilitiesTotal)
	// fact, err := factPossibilitiesFusion2(factPossibilitiesTotalTmp)
	fact, err := factPossibilitiesFusion3(factPossibilitiesTotal)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	} else {
		fmt.Printf("Result:\n")
		factManager.DisplayFacts(fact)
	}
	return nil, nil
}

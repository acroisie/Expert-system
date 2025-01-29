// package algo

// import (
//     "fmt"
//     "expert-system/src/rules"
//     "expert-system/src/factManager"
// 	"expert-system/src/v"
// )

// func factPossibilitiesFusion(factPossibilities [][]factManager.Fact) ([]factManager.Fact, *v.Error) {
//     fmt.Println("FactPossibilitiesFusion")

//     mergedFacts := make([]factManager.Fact, len(factManager.FactList))
//     copy(mergedFacts, factManager.FactList)

//     if len(factPossibilities) <= 0 {
//         return mergedFacts, &v.Error{Type: v.SOLVING, Message: "No fact possibilities to merge"}
//     }

//     for i, fact := range mergedFacts {
//         letter := fact.Letter

//         allValues := mafunc factPossibilitiesFusion(factPossibilities [][]factManager.Fact) ([]factManager.Fact, *v.Error) {
// 			fmt.Println("FactPossibilitiesFusion")
		
// 			mergedFacts := make([]factManager.Fact, len(factManager.FactList))
// 			copy(mergedFacts, factManager.FactList)
		
// 			if len(factPossibilities) <= 0 {
// 				return mergedFacts, &v.Error{Type: v.SOLVING, Message: "No fact possibilities to merge"}
// 			}
		
// 			for i, fact := range mergedFacts {
// 				letter := fact.Letter
		
// 				allValues := make([]v.Value, 0, len(factPossibilities))
// 				for _, possibility := range factPossibilities {
// 					for _, factInPossibility := range possibility {
// 						if factInPossibility.Letter == letter {
// 							allValues = append(allValues, factInPossibility.Value)
// 							break
// 						}
// 					}
// 				}
		
// 				alwaysFalse := true
// 				alwaysTrue := true
// 				for _, val := range allValues {
// 					if val != v.FALSE {
// 						alwaysFalse = false
// 					}
// 					if val != v.TRUE {
// 						alwaysTrue = false
// 					}
// 				}
		
// 				if alwaysFalse {
// 					mergedFacts[i].Value = v.FALSE
// 				} else if alwaysTrue {
// 					mergedFacts[i].Value = v.TRUE
// 				} else {
// 					fmt.Printf("Letter %c is undetermined\n", letter)
// 					var possibilitiesWithFalseValue = [][]factManager.Fact{}
// 					var possibilitiesWithTrueValue = [][]factManager.Fact{}
// 					for _, possibility := range factPossibilities {
// 						value, err := factManager.GetLetterValue(possibility, letter)
// 						if err != nil {
// 							return mergedFacts, err
// 						}
// 						possibilityWithoutLetter := factManager.SliceALetter(possibility, letter)
// 						if value == v.FALSE {
// 							possibilitiesWithFalseValue = append(possibilitiesWithFalseValue, possibilityWithoutLetter)
// 						} else if value == v.TRUE {
// 							possibilitiesWithTrueValue = append(possibilitiesWithTrueValue, possibilityWithoutLetter)
// 						}
// 					}
// 					fmt.Printf("Possibilities with False value for letter %c: %d\n", letter, len(possibilitiesWithFalseValue))
// 					fmt.Printf("Possibilities with True value for letter %c: %d\n", letter, len(possibilitiesWithTrueValue))
// 					var possibilitiesWithTrueValueToKeep = [][]factManager.Fact{}
// 					for _, truePossibility := range possibilitiesWithTrueValue {
// 						fmt.Println("Possibility with True value:")
// 						factManager.DisplayFacts(truePossibility)
// 						toRemove := false
// 						for _, falsePossibility := range possibilitiesWithFalseValue {
// 							tmp := factManager.CompareFactLists(truePossibility, falsePossibility)
// 							if tmp {
// 								toRemove = true
// 								break
// 							}
// 						}
// 						if !toRemove {
// 							possibilitiesWithTrueValueToKeep = append(possibilitiesWithTrueValueToKeep, truePossibility)
// 						}
// 					}
// 					fmt.Printf("Possibilities with True value to keep for letter %c: %d\n", letter, len(possibilitiesWithTrueValueToKeep))
// 					if len(possibilitiesWithTrueValueToKeep) > 0 {
// 						mergedFacts[i].Value = v.UNDETERMINED
// 					} else {
// 						mergedFacts[i].Value = v.FALSE
// 					}
// 				}
// 			}
		
// 			return mergedFacts, nil
// 		}ke([]v.Value, 0, len(factPossibilities))
//         for _, possibility := range factPossibilities {
//             for _, factInPossibility := range possibility {
//                 if factInPossibility.Letter == letter {
//                     allValues = append(allValues, factInPossibility.Value)
//                     break
//                 }
//             }
//         }

//         alwaysFalse := true
//         alwaysTrue := true
//         for _, val := range allValues {
//             if val != v.FALSE {
//                 alwaysFalse = false
//             }
//             if val != v.TRUE {
//                 alwaysTrue = false
//             }
//         }

//         if alwaysFalse {
//             mergedFacts[i].Value = v.FALSE
//         } else if alwaysTrue {
//             mergedFacts[i].Value = v.TRUE
//         } else {
// 			fmt.Printf("Letter %c is undetermined\n", letter)
// 			var possibilitiesWithFalseValue = [][]factManager.Fact{}
// 			var possibilitiesWithTrueValue = [][]factManager.Fact{}
//             for _, possibility := range factPossibilities {
// 				value, err := factManager.GetLetterValue(possibility, letter)
// 				if err != nil {
// 					return mergedFacts, err
// 				}
// 				possibilityWithoutLetter := factManager.SliceALetter(possibility, letter)
// 				if value == v.FALSE {
// 					possibilitiesWithFalseValue = append(possibilitiesWithFalseValue, possibilityWithoutLetter)
// 				} else if value == v.TRUE {
// 					possibilitiesWithTrueValue = append(possibilitiesWithTrueValue, possibilityWithoutLetter)
// 				}
// 			}
// 			fmt.Printf("Possibilities with False value for letter %c: %d\n", letter, len(possibilitiesWithFalseValue))
// 			fmt.Printf("Possibilities with True value for letter %c: %d\n", letter, len(possibilitiesWithTrueValue))
// 			var possibilitiesWithTrueValueToKeep = [][]factManager.Fact{}
// 			for _, truePossibility := range possibilitiesWithTrueValue {
// 				fmt.Println("Possibility with True value:")
// 				factManager.DisplayFacts(truePossibility)
// 				toRemove := false
// 				for _, falsePossibility := range possibilitiesWithFalseValue {
// 					tmp := factManager.CompareFactLists(truePossibility, falsePossibility)
// 					if tmp {
// 						toRemove = true
// 						break
// 					}
// 				}
// 				if !toRemove {
// 					possibilitiesWithTrueValueToKeep = append(possibilitiesWithTrueValueToKeep, truePossibility)
// 				}
// 			}
// 			fmt.Printf("Possibilities with True value to keep for letter %c: %d\n", letter, len(possibilitiesWithTrueValueToKeep))
// 			if len(possibilitiesWithTrueValueToKeep) > 0 {
// 				mergedFacts[i].Value = v.UNDETERMINED
// 			} else {
// 				mergedFacts[i].Value = v.FALSE
// 			}
//         }
//     }

//     return mergedFacts, nil
// }

// func countTrueValues(factList []factManager.Fact) int {
// 	count := 0
// 	for _, fact := range factList {
// 		if fact.Value == v.TRUE {
// 			count++
// 		}
// 	}
// 	return count
// }

// func getSmallestPossibilities(factPossibilities [][]factManager.Fact) ([][]factManager.Fact) {
// 	smallestPossibilities := [][]factManager.Fact{}
// 	smallestPossibilitiesSize := 0
// 	for _, possibility := range factPossibilities {
// 		trueCounter := countTrueValues(possibility)
// 		if trueCounter < smallestPossibilitiesSize || smallestPossibilitiesSize == 0 {
// 			smallestPossibilities = [][]factManager.Fact{}
// 			smallestPossibilities = append(smallestPossibilities, possibility)
// 			smallestPossibilitiesSize = trueCounter
// 		} else if trueCounter == smallestPossibilitiesSize {
// 			smallestPossibilities = append(smallestPossibilities, possibility)
// 		}
// 	}
// 	return smallestPossibilities
// }

// func factPossibilitiesFusion2(factPossibilities [][]factManager.Fact) ([]factManager.Fact, *v.Error) {
// 	same := true
// 	for i := 0; i < len(factPossibilities); {
// 		for j := i + 1; j < len(factPossibilities); {
// 			if !factManager.CompareFactLists(factPossibilities[i], factPossibilities[j]) {
// 				same = false
// 				break
// 			}
// 			j++
// 		}
// 		if !same {
// 			i++
// 		} else {
// 			break
// 		}
// 	}
// 	if same {
// 		return factPossibilities[0], nil
// 	} else {
// 		return nil, &v.Error{Type: v.SOLVING, Message: "Multiple possibilities"}
// 	}
// }

// func SliceALetter(factList []Fact, letter rune) []Fact {
//     var newFactList []Fact
//     for _, fact := range factList {
//         if fact.Letter != letter {
//             newFactList = append(newFactList, fact)
//         }
//     }
//     return newFactList
// }
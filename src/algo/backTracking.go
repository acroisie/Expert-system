package algo

import (
    "fmt"
    "expert-system/src/rules"
    "expert-system/src/factManager"
	"expert-system/src/v"
)

func backTracking() *v.Error {
	if AlgoDisplayLogs {
		fmt.Println("\n\n---------- BACKTRACKING ----------")
	}
	
	factListSave := make([]factManager.Fact, len(factManager.FactList))
	factPossibilitiesTotal := [][]factManager.Fact{}
	copy(factListSave, factManager.FactList)
	lap := 0
	maxLap := len(factManager.GetUnknowLetters())
	for lap < maxLap {

		if AlgoDisplayLogs {
			fmt.Printf("\n---------- LAP %d ----------\n", lap)
		}
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
	
		lap++
	}

	if AlgoDisplayLogs {
		fmt.Println("\n\n---------- FACTS AFTER BACKTRACKING ----------")
		fmt.Printf("factPossibilitiesTotal len = %d\n", len(factPossibilitiesTotal))
		for i, factList := range factPossibilitiesTotal {
			fmt.Printf("\n---------- FACTS %d ----------\n", i)
			factManager.DisplayFacts(factList)
		}
	}
	fact, err := determineFactListResult(factPossibilitiesTotal)
	if err != nil {
		return err
	} else {
		factManager.FactList = fact
		return nil
	}
}

func determineFactListResult(factPossibilities [][]factManager.Fact) ([]factManager.Fact, *v.Error) {

    mergedFacts := make([]factManager.Fact, len(factManager.FactList))
    copy(mergedFacts, factManager.FactList)

    if len(factPossibilities) <= 0 {
        return mergedFacts, &v.Error{Type: v.SOLVING, Message: "No fact possibilities"}
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

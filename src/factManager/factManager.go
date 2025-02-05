package factManager

import (
	"expert-system/src/v"
	"fmt"
)

const (
	Red   = "\033[31m"
	Green = "\033[32m"
	Reset = "\033[0m"
)

type Reason struct {
	Msg string
}

type Fact struct {
	Letter  rune
	Value   v.Value
	Initial bool
	Reason  Reason
}

var FactList []Fact
var FactDisplayLogs bool = false
var FactChangeCounter int = 0

func SetFactValueByLetter(letter rune, Value v.Value, force bool) *v.Error {
	fact, err := GetFactReferenceByLetter(letter)
	if err != nil {
		return err
	}
	if fact.Value == v.UNKNOWN || force {
		oldValue := fact.Value
		fact.Value = Value
		FactChangeCounter++
		LogFact(fmt.Sprintf("%c Value changed from %s to %s", letter, oldValue, fact.Value))
	} else {
		return &v.Error{Type: v.CONTRADICTION, Message: fmt.Sprintf("Fact with letter %c already has a value", letter)}
	}
	return nil
}

func GetFactReferenceByLetter(letter rune) (*Fact, *v.Error) {
	for i := range FactList {
		if FactList[i].Letter == letter {
			return &FactList[i], nil
		}
	}
	return nil, &v.Error{Type: v.FACT_NOT_FOUND, Message: fmt.Sprintf("Fact with letter %c not found", letter)}
}

func GetFactReferenceByLetterExtern(letter rune, factList []Fact) (*Fact, *v.Error) {
	for i := range factList {
		if factList[i].Letter == letter {
			return &factList[i], nil
		}
	}
	return nil, &v.Error{Type: v.FACT_NOT_FOUND, Message: fmt.Sprintf("Fact with letter %c not found", letter)}
}

func GetUnknowLetters() []rune {
	var unknowLetters []rune
	for _, fact := range FactList {
		if fact.Value == v.UNKNOWN {
			unknowLetters = append(unknowLetters, fact.Letter)
		}
	}
	return unknowLetters
}

func CompareFactLists(factList1 []Fact, factList2 []Fact) bool {
	if len(factList1) != len(factList2) {
		return false
	}
	for i := range factList1 {
		for j := range factList2 {
			if factList1[i].Letter == factList2[j].Letter && factList1[i].Value != factList2[j].Value {
				return false
			}
		}
	}
	return true
}

func SortFactListByAlphabet(factList []Fact) {
	for i := range factList {
		for j := range FactList {
			if FactList[i].Letter < FactList[j].Letter {
				FactList[i], FactList[j] = FactList[j], FactList[i]
			}
		}
	}
}

// DISPLAY

func (f Fact) String() string {
	return fmt.Sprintf("%c = %s", f.Letter, f.Value)
}

func LogFact(msg string) {
	if FactDisplayLogs {
		fmt.Println(Green, fmt.Sprintf("Fact - %s", msg), Reset)
	}
}

func DisplayFacts(facts []Fact) {
	for i, fact := range facts {
		fmt.Printf("%d: %s\n", i, fact)
	}
}

func Display() {
	DisplayFacts(FactList)
}

func DisplayFactsOneLine(facts []Fact) {
	factString := ""
	for i, fact := range FactList {
		factString += fmt.Sprintf("%s", fact)
		if i < len(FactList)-1 {
			factString += ", "
		}
	}
	fmt.Println(factString)
}

func DisplayRunesTab(runes []rune) {
	runeString := ""
	for i := 0; i < len(runes); {
		runeString += fmt.Sprintf("%c", runes[i])
		i++
	}
	fmt.Println(runeString)
}

// OLD

// func SliceALetter(factList []Fact, letter rune) []Fact {
//     var newFactList []Fact
//     for _, fact := range factList {
//         if fact.Letter != letter {
//             newFactList = append(newFactList, fact)
//         }
//     }
//     return newFactList
// }

// func GetLetterValue(factList []Fact, letter rune) (v.Value, *v.Error) {
//     for _, fact := range factList {
//         if fact.Letter == letter {
//             return fact.Value, nil
//         }
//     }
//     return v.UNKNOWN, &v.Error{Type: v.FACT_NOT_FOUND, Message: fmt.Sprintf("Fact with letter %c not found", letter)}
// }

// func RemoveElement(factPossibilities [][]Fact, index int) [][]Fact {
//     if index < 0 || index >= len(factPossibilities) {
//         return factPossibilities
//     }
//     if index == len(factPossibilities) - 1 {
//         return factPossibilities[:index]
//     }
//     return append(factPossibilities[:index], factPossibilities[index+1:]...)
// }

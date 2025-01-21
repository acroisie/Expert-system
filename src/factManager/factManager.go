package factManager

import (
	"errors"
	"expert-system/src/v"
	"fmt"
)

const (
	Red    = "\033[31m"
    Green  = "\033[32m"
	Reset  = "\033[0m"
)

type Reason struct {
    Msg string
}

type Fact struct {
    Letter rune
    Value v.Value
    Initial bool
    Reason Reason
}

var FactList []Fact
var FactDisplayLogs bool = false
var FactChangeCounter int = 0

func InitializeFactList(newFactList []Fact) {
	FactList = newFactList
}

func GetFactReferenceByLetter(letter rune) (*Fact, error) {
	for i := range FactList {
		if FactList[i].Letter == letter {
			return &FactList[i], nil
		}
	}
	return nil, errors.New(fmt.Sprintf("Fact with letter %c not found", letter))
}

func SetFactValueByLetter(letter rune, Value v.Value) error {
    fact, err := GetFactReferenceByLetter(letter)
    if err != nil {
        return err
    }
    if fact.Value != Value {
        oldValue := fact.Value
        fact.Value = Value
        FactChangeCounter++
        LogFact(fmt.Sprintf("%c Value changed from %s to %s", letter, oldValue, fact.Value))
    }
    return nil
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
    fmt.Println("---------- FACTS ----------")
    for i, fact := range facts {
        fmt.Printf("%d: %s\n", i, fact)
    }
}

func Display() {
    DisplayFacts(FactList)
}

package models

import (
	"fmt"
    "errors"
)

type Reason struct {
    Msg string
}

type Fact struct {
    Letter rune
    Value Value
    Initial bool
    Reason Reason
}

func GetFactReferenceByLetter(facts []Fact, letter rune) (*Fact, error) {
	for i := range facts {
		if facts[i].Letter == letter {
			return &facts[i], nil
		}
	}
	return nil, errors.New(fmt.Sprintf("Fact with letter %c not found", letter))
}

func SetFactValueByLetter(facts []Fact, letter rune, value Value, display bool) error {
    fact, err := GetFactReferenceByLetter(facts, letter)
    if err != nil {
        return err
    }
    oldValue := fact.Value
    fact.Value = value
    if display {
        fmt.Println(fmt.Sprintf("Fact %c value changed from %s to %s", letter, oldValue, value))
    }
    return nil
}

// DISPLAY

func (f Fact) String() string {
    return f.getFact()
}

func (f Fact) getFact() string {
    return fmt.Sprintf("%c = %s", f.Letter, f.Value)
}

func DisplayFacts(facts []Fact) {
    fmt.Println("---------- FACTS ----------")
    for i, fact := range facts {
        fmt.Printf("%d: %s\n", i, fact.getFact())
    }
}

func GetFactsMock() []Fact {
	facts := []Fact{
		{Letter: 'A', Value: TRUE, Initial: true, Reason: Reason{Msg: "Initial fact"}},
		{Letter: 'B', Value: TRUE, Initial: true, Reason: Reason{Msg: "Initial fact"}},
		{Letter: 'G', Value: TRUE, Initial: true, Reason: Reason{Msg: "Initial fact"}},
        {Letter: 'C', Value: INIT_FALSE, Initial: false, Reason: Reason{Msg: ""}},
        {Letter: 'E', Value: INIT_FALSE, Initial: false, Reason: Reason{Msg: ""}},
        {Letter: 'D', Value: INIT_FALSE, Initial: false, Reason: Reason{Msg: ""}},
        {Letter: 'F', Value: INIT_FALSE, Initial: false, Reason: Reason{Msg: ""}},
        {Letter: 'G', Value: INIT_FALSE, Initial: false, Reason: Reason{Msg: ""}},
        {Letter: 'H', Value: INIT_FALSE, Initial: false, Reason: Reason{Msg: ""}},
        {Letter: 'V', Value: INIT_FALSE, Initial: false, Reason: Reason{Msg: ""}},
        {Letter: 'W', Value: INIT_FALSE, Initial: false, Reason: Reason{Msg: ""}},
        {Letter: 'X', Value: INIT_FALSE, Initial: false, Reason: Reason{Msg: ""}},
        {Letter: 'Y', Value: INIT_FALSE, Initial: false, Reason: Reason{Msg: ""}},
        {Letter: 'Z', Value: INIT_FALSE, Initial: false, Reason: Reason{Msg: ""}},
	}
	return facts
}

package rules

import (
	"errors"
	"fmt"
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

func SetFactValueByLetter(facts []Fact, variable Variable, Value Value, display bool) error {
    fact, err := GetFactReferenceByLetter(facts, variable.Letter)
    if err != nil {
        return err
    }
    if fact.Value != Value {
        oldValue := fact.Value
        if (variable.Not) {
            fact.Value = Value.NOT()
        } else {
            fact.Value = Value
        }
        if display {
            fmt.Println(fmt.Sprintf("----- Fact %c Value changed from %s to %s -----", variable.Letter, oldValue, fact.Value))
        }
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
		{Letter: 'A', Value: v.TRUE, Initial: true, Reason: Reason{Msg: "Initial fact"}},
		{Letter: 'B', Value: v.TRUE, Initial: true, Reason: Reason{Msg: "Initial fact"}},
		{Letter: 'G', Value: v.TRUE, Initial: true, Reason: Reason{Msg: "Initial fact"}},
        {Letter: 'C', Value: v.UNKNOWN, Initial: false, Reason: Reason{Msg: ""}},
        {Letter: 'E', Value: v.UNKNOWN, Initial: false, Reason: Reason{Msg: ""}},
        {Letter: 'D', Value: v.UNKNOWN, Initial: false, Reason: Reason{Msg: ""}},
        {Letter: 'F', Value: v.UNKNOWN, Initial: false, Reason: Reason{Msg: ""}},
        {Letter: 'G', Value: v.UNKNOWN, Initial: false, Reason: Reason{Msg: ""}},
        {Letter: 'H', Value: v.UNKNOWN, Initial: false, Reason: Reason{Msg: ""}},
        {Letter: 'V', Value: v.UNKNOWN, Initial: false, Reason: Reason{Msg: ""}},
        {Letter: 'W', Value: v.UNKNOWN, Initial: false, Reason: Reason{Msg: ""}},
        {Letter: 'X', Value: v.UNKNOWN, Initial: false, Reason: Reason{Msg: ""}},
        {Letter: 'Y', Value: v.UNKNOWN, Initial: false, Reason: Reason{Msg: ""}},
        {Letter: 'Z', Value: v.UNKNOWN, Initial: false, Reason: Reason{Msg: ""}},
	}
	return facts
}

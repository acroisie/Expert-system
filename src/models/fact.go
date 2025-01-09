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



// DISPLAY

func (f Fact) String() string {
    return f.getFact()
}

func (f Fact) getFact() string {
    return fmt.Sprintf("%c = %t", f.Letter, f.Value)
}

func DisplayFacts(facts []Fact) {
    for i, fact := range facts {
        fmt.Printf("%d: %s\n", i+1, fact.getFact())
    }
}

func GetFactsMock() []Fact {
	facts := []Fact{
		{Letter: 'A', Value: TRUE, Initial: true, Reason: Reason{Msg: "Initial fact"}},
		{Letter: 'B', Value: TRUE, Initial: true, Reason: Reason{Msg: "Initial fact"}},
		{Letter: 'G', Value: TRUE, Initial: true, Reason: Reason{Msg: "Initial fact"}},
	}
	return facts
}

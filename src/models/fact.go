package models

import (
	"fmt"
)

type Reason struct {
    Msg string
}

type Fact struct {
    Letter rune
    Value bool
    Initial bool
    Reason Reason
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
		{Letter: 'A', Value: true, Initial: true, Reason: Reason{Msg: "Initial fact"}},
		{Letter: 'B', Value: true, Initial: true, Reason: Reason{Msg: "Initial fact"}},
		{Letter: 'G', Value: true, Initial: true, Reason: Reason{Msg: "Initial fact"}},
	}
	return facts
}

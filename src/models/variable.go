package models

import (
	"errors"
)

type Variable struct {
    Letter rune
	Not bool
}

func (v Variable) GetValueByFacts(facts []Fact) (Value, error) {
	for i := range facts {
		if facts[i].Letter == v.Letter {
			return v.valueByFact(facts[i]), nil
		}
	}
	return UNDETERMINED, errors.New("Fact not found")
}


func (v Variable) valueByFact(fact Fact) Value {
	if (v.Not) {
		return fact.Value.NOT()
	} else {
		return fact.Value
	}
}


// DISPLAY

func (v Variable) String() string {
	if (v.Not) {
		return "!" + string(v.Letter)
	} else {
		return string(v.Letter)
	}
}

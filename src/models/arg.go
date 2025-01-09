package models

import (
	"errors"
)

type Arg struct {
    Op LogicalOperator
    Letter rune
	Not bool
}

func (c Arg) value(fact Fact) bool {
	if (c.Not) {
		return !fact.Value
	} else {
		return fact.Value
	}
}

func (c Arg) solving(previousValue Value, facts []Fact) (bool, error) {
	fact, err := GetFactReferenceByLetter(facts, c.Letter)
	if err != nil {
		return false, err
	}

	if previousLetter == 0 {
		return c.value(*fact), nil
	}

	if !c.Op.isValid() {
		return false, errors.New(fmt.Sprintf("Unknown operator: %s", c.Op.toString()))
	}
	

	return result, nil
}



// DISPLAY

func (c Arg) getLetter() string {
	if (c.Not) {
		return "!" + string(c.Letter)
	} else {
		return string(c.Letter)
	}
}

package models

type Variable struct {
    Letter rune
	Not bool
}

func (v Variable) value(fact Fact) Value {
	if (v.Not) {
		return fact.Value.NOT()
	} else {
		return fact.Value
	}
}

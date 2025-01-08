package models

type Arg struct {
    Op LogicalOperator
    Letter rune
	Not bool
}

func (c Arg) getLetter() string {
	if (c.Not) {
		return "!" + string(c.Letter)
	} else {
		return string(c.Letter)
	}
}
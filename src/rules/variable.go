package rules

type Variable struct {
	Letter rune
	Not    bool
}

// DISPLAY

func (v Variable) String() string {
	if v.Not {
		return "!" + string(v.Letter)
	} else {
		return string(v.Letter)
	}
}

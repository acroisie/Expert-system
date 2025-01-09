package models

type Value int

const (
    INIT_FALSE Value = iota
    FALSE
    TRUE
    UNDETERMINED
)

func (v Value) normalize() Value {
	if v == INIT_FALSE {
		return FALSE
	}
	return v
}

func (firstValue Value) AND(secondValue Value) Value {
	a := firstValue.normalize()
	b := secondValue.normalize()
	switch a {
		case FALSE:
			return FALSE
		case TRUE:
			if b == TRUE {
				return TRUE
			} else if b == FALSE {
				return FALSE
			} else {
				return UNDETERMINED
			}
		case UNDETERMINED:
			if b == TRUE {
				return UNDETERMINED
			} else if b == FALSE {
				return FALSE
			} else {
				return UNDETERMINED
			}
		default:
			return UNDETERMINED
	}
}

func (firstValue Value) OR(secondValue Value) Value {
	a := firstValue.normalize()
	b := secondValue.normalize()
	switch a {
		case FALSE:
			if b == TRUE {
				return TRUE
			} else if b == FALSE {
				return FALSE
			} else {
				return UNDETERMINED
			}
		case TRUE:
			return TRUE
		case UNDETERMINED:
			if b == TRUE {
				return TRUE
			} else if b == FALSE {
				return UNDETERMINED
			} else {
				return UNDETERMINED
			}
		default:
			return UNDETERMINED
	}
}

func (firstValue Value) XOR(secondValue Value) Value {
	a := firstValue.normalize()
	b := secondValue.normalize()
	switch a {
		case FALSE:
			if b == TRUE {
				return TRUE
			} else if b == FALSE {
				return FALSE
			} else {
				return UNDETERMINED
			}
		case TRUE:
			if b == TRUE {
				return FALSE
			} else if b == FALSE {
				return TRUE
			} else {
				return UNDETERMINED
			}
		case UNDETERMINED:
			return UNDETERMINED
		default:
			return UNDETERMINED
	}
}

func (v Value) NOT() Value {
	value := v.normalize()
	switch value {
		case FALSE:
			return TRUE
		case TRUE:
			return FALSE
		case UNDETERMINED:
			return UNDETERMINED
		default:
			return UNDETERMINED
	}
}

func (v Value) String() string {
	switch v {
		case FALSE:
			return "FALSE"
		case TRUE:
			return "TRUE"
		case UNDETERMINED:
			return "UNDETERMINED"
		case INIT_FALSE:
			return "INIT_FALSE"
		default:
			return "UNKNOWN"
	}
}

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
			return FALSE, nil
		case TRUE:
			if b == TRUE {
				return TRUE, nil
			} else if b == FALSE {
				return FALSE, nil
			} else {
				return UNDETERMINED, nil
			}
		case UNDETERMINED:
			if b == TRUE {
				return UNDETERMINED, nil
			} else if b == FALSE {
				return FALSE, nil
			} else {
				return UNDETERMINED, nil
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
				return TRUE, nil
			} else if b == FALSE {
				return FALSE, nil
			} else {
				return UNDETERMINED, nil
			}
		case TRUE:
			return TRUE, nil
		case UNDETERMINED:
			if b == TRUE {
				return TRUE, nil
			} else if b == FALSE {
				return UNDETERMINED, nil
			} else {
				return UNDETERMINED, nil
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
				return TRUE, nil
			} else if b == FALSE {
				return FALSE, nil
			} else {
				return UNDETERMINED, nil
			}
		case TRUE:
			if b == TRUE {
				return FALSE, nil
			} else if b == FALSE {
				return TRUE, nil
			} else {
				return UNDETERMINED, nil
			}
		case UNDETERMINED:
			return UNDETERMINED, nil
		default:
			return UNDETERMINED
	}
}



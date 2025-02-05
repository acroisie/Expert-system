package v

type Value int

const (
	UNKNOWN Value = iota
	FALSE
	TRUE
	UNDETERMINED
)

func (v Value) normalize() Value {
	if v == UNKNOWN {
		return FALSE
	}
	return v
}

func (firstValue Value) AND(secondValue Value) Value {
	// a := firstValue.normalize()
	// b := secondValue.normalize()
	a := firstValue
	b := secondValue
	switch a {
	case FALSE:
		return FALSE
	case TRUE:
		if b == TRUE {
			return TRUE
		} else if b == FALSE {
			return FALSE
		} else if b == UNKNOWN {
			return UNKNOWN
		} else {
			return UNDETERMINED
		}
	case UNKNOWN:
		if b == TRUE {
			return UNKNOWN
		} else if b == FALSE {
			return FALSE
		} else if b == UNKNOWN {
			return UNKNOWN
		} else {
			return UNDETERMINED
		}
	case UNDETERMINED:
		if b == TRUE {
			return UNDETERMINED
		} else if b == FALSE {
			return FALSE
		} else if b == UNKNOWN {
			return UNDETERMINED
		} else {
			return UNDETERMINED
		}
	default:
		return UNDETERMINED
	}
}

func (firstValue Value) OR(secondValue Value) Value {
	// a := firstValue.normalize()
	// b := secondValue.normalize()
	a := firstValue
	b := secondValue
	switch a {
	case FALSE:
		if b == TRUE {
			return TRUE
		} else if b == FALSE {
			return FALSE
		} else if b == UNKNOWN {
			return UNKNOWN
		} else {
			return UNDETERMINED
		}
	case TRUE:
		return TRUE
	case UNKNOWN:
		if b == TRUE {
			return TRUE
		} else if b == FALSE {
			return UNKNOWN
		} else if b == UNKNOWN {
			return UNKNOWN
		} else {
			return UNDETERMINED
		}
	case UNDETERMINED:
		if b == TRUE {
			return TRUE
		} else if b == FALSE {
			return UNDETERMINED
		} else if b == UNKNOWN {
			return UNDETERMINED
		} else {
			return UNDETERMINED
		}
	default:
		return UNDETERMINED
	}
}

func (firstValue Value) XOR(secondValue Value) Value {
	// a := firstValue.normalize()
	// b := secondValue.normalize()
	a := firstValue
	b := secondValue
	switch a {
	case FALSE:
		if b == TRUE {
			return TRUE
		} else if b == FALSE {
			return FALSE
		} else if b == UNKNOWN {
			return UNKNOWN
		} else {
			return UNDETERMINED
		}
	case TRUE:
		if b == TRUE {
			return FALSE
		} else if b == FALSE {
			return TRUE
		} else if b == UNKNOWN {
			return UNKNOWN
		} else {
			return UNDETERMINED
		}
	case UNKNOWN:
		if b == TRUE {
			return UNKNOWN
		} else if b == FALSE {
			return UNKNOWN
		} else if b == UNKNOWN {
			return UNKNOWN
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
	// value := v.normalize()
	value := v
	switch value {
	case FALSE:
		return TRUE
	case TRUE:
		return FALSE
	case UNKNOWN:
		return UNKNOWN
	case UNDETERMINED:
		return UNDETERMINED
	default:
		return UNDETERMINED
	}
}

func (v Value) Real() bool {
	return v == TRUE || v == FALSE
}

func (res Value) FindUnknown_OR(know Value) Value {
	if res == TRUE {
		if know == TRUE {
			return UNKNOWN
		} else {
			return TRUE
		}
	} else {
		return FALSE
	}
}

func (res Value) FindUnknown_AND(know Value) Value {
	if res == TRUE {
		return know
	} else {
		return FALSE
	}
}

func (res Value) FindUnknown_XOR(know Value) Value {
	if res == TRUE {
		if know == TRUE {
			return FALSE
		} else {
			return TRUE
		}
	} else {
		return know
	}
}

func (res Value) FindTwoUnknown_OR() (Value, Value) {
	if res == TRUE {
		return UNKNOWN, UNKNOWN
	} else {
		return UNKNOWN, UNKNOWN
	}
}

func (res Value) FindTwoUnknown_AND() (Value, Value) {
	if res == TRUE {
		return TRUE, TRUE
	} else {
		return UNKNOWN, UNKNOWN
	}
}

func (res Value) FindTwoUnknown_XOR() (Value, Value) {
	if res == TRUE {
		return UNKNOWN, UNKNOWN
	} else {
		return UNKNOWN, UNKNOWN
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
	case UNKNOWN:
		return "UNKNOWN"
	default:
		return "UNKNOWN"
	}
}

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
			} else if b == INIT_FALSE {
				return INIT_FALSE
			} else {
				return UNDETERMINED
			}
		case INIT_FALSE:
			if b == TRUE {
				return INIT_FALSE
			} else if b == FALSE {
				return FALSE
			} else if b == INIT_FALSE {
				return INIT_FALSE
			} else {
				return UNDETERMINED
			}
		case UNDETERMINED:
			if b == TRUE {
				return UNDETERMINED
			} else if b == FALSE {
				return FALSE
			} else if b == INIT_FALSE {
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
			} else if b == INIT_FALSE {
				return INIT_FALSE
			} else {
				return UNDETERMINED
			}
		case TRUE:
			return TRUE
		case INIT_FALSE:
			if b == TRUE {
				return TRUE
			} else if b == FALSE {
				return INIT_FALSE
			} else if b == INIT_FALSE {
				return INIT_FALSE
			} else {
				return UNDETERMINED
			}
		case UNDETERMINED:
			if b == TRUE {
				return TRUE
			} else if b == FALSE {
				return UNDETERMINED
			} else if b == INIT_FALSE {
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
			} else if b == INIT_FALSE {
				return INIT_FALSE
			} else {
				return UNDETERMINED
			}
		case TRUE:
			if b == TRUE {
				return FALSE
			} else if b == FALSE {
				return TRUE
			} else if b == INIT_FALSE {
				return INIT_FALSE
			} else {
				return UNDETERMINED
			}
		case INIT_FALSE:
			if b == TRUE {
				return INIT_FALSE
			} else if b == FALSE {
				return INIT_FALSE
			} else if b == INIT_FALSE {
				return INIT_FALSE
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
		case INIT_FALSE:
			return INIT_FALSE
		case UNDETERMINED:
			return UNDETERMINED
		default:
			return UNDETERMINED
	}
}

func (v Value) real() bool {
	return v == TRUE || v == FALSE
}

func (res Value) findUnknown_OR(know Value) Value {
	if res == TRUE {
		if know == TRUE {
			return INIT_FALSE
		} else {
			return TRUE
		}
	} else  {
		return FALSE
	}
}

func (res Value) findUnknown_AND(know Value) Value {
	if res == TRUE {
		return know
	} else {
		return FALSE
	}
}

func (res Value) findUnknown_XOR(know Value) Value {
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

// func (res Value) findUnknown_Check(A Variable, B Variable) (bool, error) {
// 	if !res.real() {
// 		return false, A, errors.New("Le résultat doit être connu")
// 	}
// 	if A.Value.real() && B.Value.real() {
// 		if res == A.Value.AND(B.Value) {
// 			return false, A, nil
// 		} else {
// 			return false, A, errors.New("CONTRADICTION Les valeurs de A et B ne correspondent pas au résultat")
// 		}
// 	}
// 	if A.Value == INIT_FALSE && B.Value == INIT_FALSE {
// 		return false, nil
// 	}
// 	return true, nil
// }

// func (res Value) findUnknown_AND(A Variable, B Variable) (bool, Variable, error) {
// 	check, err := res.findUnknown_Check(A, B)
// 	if !check {
// 		return false, A, err
// 	}
// 	if A.Value == INIT_FALSE {
// 		return true, Variable{Letter: A.Letter, Value: res}, nil
// 	} else if B.Value == INIT_FALSE {
// 		return true, Variable{Letter: B.Letter, Value: res}, nil
// 	}
// 	return false, Variable{}, errors.New("CONTRADICTION : Aucune variable n'est inconnue")
// }

// func (res Value) FindUnknown_OR(A Variable, B Variable) (bool, Variable, error) {
// 	check, err := res.findUnknown_Check(A, B)
// 	if !check {
// 		return false, A, err
// 	}

// 	if res == TRUE {
// 		if A.Value == INIT_FALSE {
// 			return true, Variable{Letter: A.Letter, Value: TRUE}, nil
// 		} else if B.Value == INIT_FALSE {
// 			return true, Variable{Letter: B.Letter, Value: TRUE}, nil
// 		}
// 	} else if res == FALSE {
// 		if A.Value == INIT_FALSE {
// 			return true, Variable{Letter: A.Letter, Value: FALSE}, nil
// 		} else if B.Value == INIT_FALSE {
// 			return true, Variable{Letter: B.Letter, Value: FALSE}, nil
// 		}
// 	}
// 	return false, Variable{}, errors.New("CONTRADICTION : Aucune variable n'est inconnue")
// }

// func (res Value) FindUnknown_XOR(A Variable, B Variable) (bool, Variable, error) {
// 	check, err := res.findUnknown_Check(A, B)
// 	if !check {
// 		return false, A, err
// 	}

// 	if res == TRUE {
// 		if A.Value == INIT_FALSE {
// 			return true, Variable{Letter: A.Letter, Value: !B.Value.real()}, nil
// 		} else if B.Value == INIT_FALSE {
// 			return true, Variable{Letter: B.Letter, Value: !A.Value.real()}, nil
// 		}
// 	} else if res == FALSE {
// 		if A.Value == INIT_FALSE {
// 			return true, Variable{Letter: A.Letter, Value: B.Value.real()}, nil
// 		} else if B.Value == INIT_FALSE {
// 			return true, Variable{Letter: B.Letter, Value: A.Value.real()}, nil
// 		}
// 	}
// 	return false, Variable{}, errors.New("CONTRADICTION : Aucune variable n'est inconnue")
// }

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

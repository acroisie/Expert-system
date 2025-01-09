package models

import (
    "fmt"
    "errors"
)

type ConditionalOperator int

const (
    IMPLIES ConditionalOperator = iota
    IFF
)

func (op ConditionalOperator) solve(a Value, b Value) (Value, Value, error) {
    fmt.Sprintf("ConditionalOperator solving : %s %s %s", a, op, b)
    if !op.isValid() {
        return UNDETERMINED, UNDETERMINED, errors.New(fmt.Sprintf("Unknown operator: %s", op.toString()))
    }
    switch a {
        case FALSE:
            if b == TRUE {
                return FALSE, TRUE, errors.New("CONTRADICTION : FALSE => TRUE")
            } else if b == FALSE {
                return FALSE, FALSE, nil
            } else if b == INIT_FALSE {
                return FALSE, FALSE, nil
            } else {
                return FALSE, FALSE, nil
            }
        case TRUE:
            if b == TRUE {
                return TRUE, TRUE, nil
            } else if b == FALSE {
                return TRUE, FALSE, errors.New("CONTRADICTION : TRUE => FALSE")
            } else if b == INIT_FALSE {
                return TRUE, TRUE, nil
            } else {
                return TRUE, TRUE, nil
            }
        case INIT_FALSE:
            return a, b, nil
        case UNDETERMINED:
            return UNDETERMINED, b, nil
        default:
            return UNDETERMINED, UNDETERMINED, errors.New(fmt.Sprintf("Unknown value: %s", a))
    }
}

func (op ConditionalOperator) isValid() bool {
    return op == IMPLIES
}

// DISPLAY

func (op ConditionalOperator) toString() string {
    return [...]string{"IMPLIES", "IFF"}[op]
}

func (op ConditionalOperator) getSymbol() string {
    return [...]string{"=>", "<=>"}[op]
}

func (op ConditionalOperator) String() string {
	return op.getSymbol()
}

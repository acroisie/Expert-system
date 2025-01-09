package models

import (
	"errors"
	"fmt"
)

type LogicalOperator int

const (
    NOTHING LogicalOperator = iota
    AND
    OR
    XOR
)

func (op LogicalOperator) isValid() bool {
    return op > NOTHING && op <= XOR
}

func (op LogicalOperator) solve(a Value, b Value, display bool) (Value, error) {
    if !op.isValid() {
        return UNDETERMINED, errors.New(fmt.Sprintf("Unknown operator: %s", op.toString()))
    }
    if display {
        fmt.Println(fmt.Sprintf("LogicalOperator solving : %s %s %s", a, op, b))
    }
    switch op {
        case AND:
            return a.AND(b), nil
        case OR:
            return a.OR(b), nil
        case XOR:
            return a.XOR(b), nil
        default:
            return UNDETERMINED, errors.New(fmt.Sprintf("Unknown operator: %s", op.toString()))
    }
}

// DISPLAY

func (op LogicalOperator) toString() string {
    return [...]string{"NOTHING", "AND", "OR", "XOR"}[op]
}

func (op LogicalOperator) getSymbol() string {
    return [...]string{"", "+", "|", "^"}[op]
}

func (op LogicalOperator) String() string {
	return op.getSymbol()
}

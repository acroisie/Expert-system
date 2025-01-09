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

func (op LogicalOperator) solve(a Value, b Value) (Value, error) {
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

func (op LogicalOperator) toString() string {
    return [...]string{"NOTHING", "AND", "OR", "XOR"}[op]
}

func (op LogicalOperator) getSymbol() string {
    return [...]string{"", "+", "|", "^"}[op]
}

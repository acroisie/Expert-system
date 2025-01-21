package rules

import (
	"errors"
	"expert-system/src/v"
	"fmt"
)

var LogicalOperatorDisplayLogs bool = false

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

func (op LogicalOperator) solve(a v.Value, b v.Value) (v.Value, error) {
    if !op.isValid() {
        return v.UNDETERMINED, errors.New(fmt.Sprintf("Unknown operator: %s", op.toString()))
    }
    var result v.Value
    switch op {
        case AND:
            result = a.AND(b)
        case OR:
            result = a.OR(b)
        case XOR:
            result = a.XOR(b)
        default:
            return v.UNDETERMINED, errors.New(fmt.Sprintf("Unknown operator: %s", op.toString()))
    }
    LogLogicalOp(fmt.Sprintf("Solving : %s %s %s", a, op, b))
    return result, nil
}

// DISPLAY

func LogLogicalOp(msg string) {
    if LogicalOperatorDisplayLogs {
        fmt.Println(fmt.Sprintf("LogicalOperator - %s", msg))
    }
}

func (op LogicalOperator) toString() string {
    return [...]string{"NOTHING", "AND", "OR", "XOR"}[op]
}

func (op LogicalOperator) getSymbol() string {
    return [...]string{"", "+", "|", "^"}[op]
}

func (op LogicalOperator) String() string {
	return op.getSymbol()
}

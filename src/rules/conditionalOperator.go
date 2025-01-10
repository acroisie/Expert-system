package rules

import (
    "fmt"
    "errors"
    "expert/v"
)

var ConditionalOperatorDisplayLogs bool = false

type ConditionalOperator int

const (
    IMPLIES ConditionalOperator = iota
    IFF
)

func (op ConditionalOperator) solve(a v.Value, b v.Value) (v.Value, v.Value, error) {
    LogConditionalOp(fmt.Sprintf("Solving : %s %s %s", a, op, b))
    if !op.isValid() {
        return v.UNDETERMINED, v.UNDETERMINED, errors.New(fmt.Sprintf("Unknown operator: %s", op.toString()))
    }
    switch a {
        case v.FALSE:
            if b == v.TRUE {
                return v.FALSE, v.TRUE, errors.New("CONTRADICTION : v.FALSE => v.TRUE")
            } else if b == v.FALSE {
                return v.FALSE, v.FALSE, nil
            } else if b == v.UNKNOW {
                return v.FALSE, v.FALSE, nil
            } else {
                return v.FALSE, v.FALSE, nil
            }
        case v.TRUE:
            if b == v.TRUE {
                return v.TRUE, v.TRUE, nil
            } else if b == v.FALSE {
                return v.TRUE, v.FALSE, errors.New("CONTRADICTION : v.TRUE => v.FALSE")
            } else if b == v.UNKNOW {
                return v.TRUE, v.TRUE, nil
            } else {
                return v.TRUE, v.TRUE, nil
            }
        case v.UNKNOW:
            return a, b, nil
        case v.UNDETERMINED:
            return v.UNDETERMINED, b, nil
        default:
            return v.UNDETERMINED, v.UNDETERMINED, errors.New(fmt.Sprintf("Unknown Value: %s", a))
    }
}

func (op ConditionalOperator) isValid() bool {
    return op == IMPLIES
}

// DISPLAY

func LogConditionalOp(msg string) {
    if ConditionalOperatorDisplayLogs {
        fmt.Println(fmt.Sprintf("ConditionalOperator - %s", msg))
    }
}

func (op ConditionalOperator) toString() string {
    return [...]string{"IMPLIES", "IFF"}[op]
}

func (op ConditionalOperator) getSymbol() string {
    return [...]string{"=>", "<=>"}[op]
}

func (op ConditionalOperator) String() string {
	return op.getSymbol()
}

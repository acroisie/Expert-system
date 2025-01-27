package rules

import (
	"expert-system/src/v"
	"fmt"
)

var ConditionalOperatorDisplayLogs bool = false

type ConditionalOperator int

const (
	IMPLIES ConditionalOperator = iota
	IFF
)

func (op ConditionalOperator) solve(a v.Value, b v.Value) (v.Value, v.Value, *v.Error) {
    LogConditionalOp(fmt.Sprintf("Solving : %s %s %s", a, op, b))
    if !op.isValid() {
        return v.UNDETERMINED, v.UNDETERMINED, &v.Error{Type: v.SOLVING, Message: fmt.Sprintf("Invalid operator : %s", op)}
    }
    switch a {
        case v.FALSE:
            if b == v.TRUE {
                return v.FALSE, v.TRUE, &v.Error{Type: v.CONTRADICTION, Message: fmt.Sprintf("%s %s %s", a, op, b)}
            } else if b == v.FALSE {
                return v.FALSE, v.FALSE, nil
            } else if b == v.UNKNOWN {
                return v.FALSE, v.FALSE, nil
            } else {
                return v.FALSE, v.FALSE, nil
            }
        case v.TRUE:
            if b == v.TRUE {
                return v.TRUE, v.TRUE, nil
            } else if b == v.FALSE {
                return v.TRUE, v.FALSE, &v.Error{Type: v.CONTRADICTION, Message: fmt.Sprintf("%s %s %s", a, op, b)}
            } else if b == v.UNKNOWN {
                return v.TRUE, v.TRUE, nil
            } else {
                return v.TRUE, v.TRUE, nil
            }
        case v.UNKNOWN:
            return a, b, nil
        case v.UNDETERMINED:
            return v.UNDETERMINED, b, nil
        default:
            return v.UNDETERMINED, v.UNDETERMINED, &v.Error{Type: v.SOLVING, Message: fmt.Sprintf("Invalid value : %s", a)}
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

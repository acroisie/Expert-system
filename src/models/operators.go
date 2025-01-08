package models

type LogicalOperator int

const (
    NOTHING LogicalOperator = iota
    AND
    OR
    XOR
)

func (op LogicalOperator) toString() string {
    return [...]string{"NOTHING", "AND", "OR", "XOR"}[op]
}

func (op LogicalOperator) getSymbol() string {
    return [...]string{"", "+", "|", "^"}[op]
}

type ConditionalOperator int

const (
    IMPLIES ConditionalOperator = iota
    IFF
)

func (op ConditionalOperator) toString() string {
    return [...]string{"IMPLIES", "IFF"}[op]
}

func (op ConditionalOperator) getSymbol() string {
    return [...]string{"=>", "<=>"}[op]
}

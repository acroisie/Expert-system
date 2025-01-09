package models

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

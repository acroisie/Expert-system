package models

type Rule struct {
	Op ConditionalOperator
	LeftExpressionGroup *ExpressionGroup
	RightExpressionGroup *ExpressionGroup
	LeftVariable *Variable
	RightVariable *Variable
}
package parser

import (
	"expert-system/src/rules"
	"fmt"
)

type Parser struct {
	lexer   *Lexer
	currTok Token
}

func NewParser(input string) *Parser {
	p := &Parser{lexer: NewLexer(input)}
	p.nextToken()
	return p
}

func (p *Parser) nextToken() {
	p.currTok = p.lexer.NextToken()
}

func (p *Parser) ParseRule() (*rules.Rule, error) {
	leftExpr, err := p.parseExpression()
	if err != nil {
		return nil, err
	}

	if p.currTok.Type != TKN_IMPLIES && p.currTok.Type != TKN_IFF {
		return nil, fmt.Errorf("expected '=>' or '<=>' but got %v", p.currTok)
	}

	var op rules.ConditionalOperator
	if p.currTok.Type == TKN_IMPLIES {
		op = rules.IMPLIES
	} else {
		op = rules.IFF
	}

	p.nextToken()

	rightExpr, err := p.parseExpression()
	if err != nil {
		return nil, err
	}

	leftExpr = p.cleanExpressionGroup(leftExpr)
	rightExpr = p.cleanExpressionGroup(rightExpr)

	return &rules.Rule{
		Op:                   op,
		LeftExpressionGroup:  leftExpr,
		RightExpressionGroup: rightExpr,
	}, nil
}

func (p *Parser) cleanExpressionGroup(eg *rules.ExpressionGroup) *rules.ExpressionGroup {
	if eg == nil {
		return nil
	}

	if eg.Op == rules.NOTHING && eg.LeftVariable != nil && eg.RightVariable == nil &&
		eg.LeftExpressionGroup == nil && eg.RightExpressionGroup == nil {
		return &rules.ExpressionGroup{
			Op:           rules.NOTHING,
			LeftVariable: eg.LeftVariable,
		}
	}

	if eg.LeftExpressionGroup != nil {
		eg.LeftExpressionGroup = p.cleanExpressionGroup(eg.LeftExpressionGroup)
	}
	if eg.RightExpressionGroup != nil {
		eg.RightExpressionGroup = p.cleanExpressionGroup(eg.RightExpressionGroup)
	}

	return eg
}

func (p *Parser) parseExpression() (*rules.ExpressionGroup, error) {
	leftExpr, err := p.parseTerm()
	if err != nil {
		return nil, err
	}

	for p.currTok.Type == TKN_OR || p.currTok.Type == TKN_XOR {
		opToken := p.currTok
		p.nextToken()

		rightExpr, err := p.parseTerm()
		if err != nil {
			return nil, err
		}

		newNode := &rules.ExpressionGroup{}
		if opToken.Type == TKN_OR {
			newNode.Op = rules.OR
		} else {
			newNode.Op = rules.XOR
		}

		newNode.LeftExpressionGroup = leftExpr
		newNode.RightExpressionGroup = rightExpr

		leftExpr = newNode
	}

	return leftExpr, nil
}

func (p *Parser) parseTerm() (*rules.ExpressionGroup, error) {
	leftExpr, err := p.parseFactor()
	if err != nil {
		return nil, err
	}

	for p.currTok.Type == TKN_AND {
		p.nextToken()

		rightExpr, err := p.parseFactor()
		if err != nil {
			return nil, err
		}

		newNode := &rules.ExpressionGroup{
			Op:                   rules.AND,
			LeftExpressionGroup:  leftExpr,
			RightExpressionGroup: rightExpr,
		}

		leftExpr = newNode
	}

	return leftExpr, nil
}

func (p *Parser) parseFactor() (*rules.ExpressionGroup, error) {
	if p.currTok.Type == TKN_NOT {
		p.nextToken()

		expr, err := p.parseFactor()
		if err != nil {
			return nil, err
		}

		if expr.Op == rules.NOTHING && expr.LeftVariable != nil && expr.RightVariable == nil {
			expr.LeftVariable.Not = !expr.LeftVariable.Not
			return expr, nil
		}

		newNode := &rules.ExpressionGroup{
			Op:                  rules.NOTHING,
			LeftVariable:        &rules.Variable{Letter: 0, Not: true},
			LeftExpressionGroup: expr,
		}

		return newNode, nil
	}

	if p.currTok.Type == TKN_LPAREN {
		p.nextToken()

		expr, err := p.parseExpression()
		if err != nil {
			return nil, err
		}

		if p.currTok.Type != TKN_RPAREN {
			return nil, fmt.Errorf("expected ')' but got %v", p.currTok)
		}

		p.nextToken()
		return expr, nil
	}

	if p.currTok.Type == TKN_VAR {
		newNode := &rules.ExpressionGroup{
			Op: rules.NOTHING,
			LeftVariable: &rules.Variable{
				Letter: rune(p.currTok.Value[0]),
				Not:    false,
			},
		}
		p.nextToken()
		return newNode, nil
	}

	return nil, fmt.Errorf("expected variable, '(', '!' but got %v", p.currTok)
}

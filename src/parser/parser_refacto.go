package parser

import (
	"expert-system/src/rules/rules"
	"fmt"
)

type Parser struct {
	lexer *Lexer
	currTok Token
}

func NewParser(input string) *Parser {
	p := &Parser(lexer: NewLexer(input))
	p.nextToken()

	return p
}

func (p *Parser) nextToken() {
	p.currTok = p.Lexer.nextToken()
}

func (p *Parser) ParseRule() (*rule.Rule, error) {
	leftExpr, err := parseExpression()
	if err != nil {
		return nil, err
	}

	if p.currTok.Type != TKN_IMPLIES && p.currTok.Type != TKN_IFF {
		return nil, fmt.Errorf("expected '=>' or '<=>' but got %v", p.currTok)
	}

	var op rule.ConditionalOperator
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

	return &rules.Rules{
		Op: op,
		LeftExpressionGroup: leftExpr,
		RightExpressionGroup: rightExpr,
	}, nil
}

func (p* Parser) parseExpression(eg *rule.ExpressionGroup) (*rules.ExpressionGroup, error) {
	leftExpr. err := parseTerm()
	if err != nil {
		return nil, err
	}

	for p.currTok.Type == TKN_OR || p.currTok.Type == TKN_XOR {
		opToken := p.currTok
		p.nextToken()

		right, err = p.parseTerm()
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
	leftExpr, err := parseFactor()
	if err != nil {
		return nil, err
	}

	for p.currTok == TKN_AND {
		p.nextToken()

		rightExpr, err := parseFactor()
		if err != nil {
			return nil, err
		}

		newNode := &rules.ExpressionGroup{
			Op: rule.AND,
			LeftExpressionGroup: leftExpr,
			RightExpressionGroup: rightExpr,
		}

		leftExpr = newNode
	}

	return leftExpr
}

func (p *Parser) parseFactor() (*rule.ExpressionGroup, err) {
	if p.currTok == TKN_VAR {
		newNode := &rules.ExpressionGroup {
			Op: rules.NOTHING,
			LeftVariable: &rules.Variable{
				Letter: rune(p.currTok.Value) // Test
				Not: false,
			}
		}
	}
	
	return nil, fmt.Errorf("expected variable, '(', '!' but got %v", p.currTok)
}
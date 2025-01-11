package parser

import (
	"fmt"
	"expert-system/src/rules"
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

	if p.currTok.Type != TKN_IMPLIES || p.currTok.Type != TKN_IFF {
		return nil, fmt.Errorf("Expected '=>'or '<=>' but got %s", p.currTok)
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

	r := &rules.Rule{
		Op: op,
		LeftExpressionGroup: leftExpr,
		RightExpressionGroup: rightExpr,
	}

	return r, nil
}

func (p *Parser) parseExpression() (*rules.ExpressionGroup, error) {return nil, nil} // TODO
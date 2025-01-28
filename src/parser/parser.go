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
		return nil, fmt.Errorf("expected '=>'or '<=>' but got %v", p.currTok)
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

    leftEG, leftVar := simplifyExpression(leftExpr)
    rightEG, rightVar := simplifyExpression(rightExpr)

	if leftEG != nil {
		leftEG = exploreEp(leftEG)
	}
	if rightEG != nil {
		rightEG = exploreEp(rightEG)
	}

    return &rules.Rule{
        Op:                   op,
        LeftExpressionGroup:  leftEG,
        RightExpressionGroup: rightEG,
        LeftVariable:         leftVar,
        RightVariable:        rightVar,
    }, nil
}

// (A AND B) AND C
func exploreEp(eg *rules.ExpressionGroup) *rules.ExpressionGroup {
	leftEg, leftVariable := simplifyExpression(eg.LeftExpressionGroup)
	rightEg, rightVariable := simplifyExpression(eg.RightExpressionGroup)

	if leftEg != nil {
		leftEg = exploreEp(leftEg)
	}
	if rightEg != nil {
		rightEg = exploreEp(rightEg)
	}

	return &rules.ExpressionGroup{
		Op: eg.Op,
		LeftExpressionGroup: leftEg,
		RightExpressionGroup: rightEg,
		LeftVariable: leftVariable,
		RightVariable: rightVariable,
	}
}

func simplifyExpression(eg *rules.ExpressionGroup) (*rules.ExpressionGroup, *rules.Variable) {
    if eg == nil {
        return nil, nil
    }

    if eg.Op == rules.NOTHING &&
       eg.LeftVariable != nil &&
       eg.RightVariable == nil &&
       eg.LeftExpressionGroup == nil &&
       eg.RightExpressionGroup == nil {
        return nil, eg.LeftVariable
    }

	if eg.Op == rules.NOTHING &&
	eg.RightVariable != nil &&
	eg.LeftVariable == nil &&
	eg.LeftExpressionGroup == nil &&
	eg.RightExpressionGroup == nil {
	 return nil, eg.RightVariable
 }

    return eg, nil
}

func (p *Parser) parseExpression() (*rules.ExpressionGroup, error) {
	left, err := p.parseTerm()
	if err != nil {
		return nil, err
	}

	for p.currTok.Type == TKN_OR || p.currTok.Type == TKN_XOR {
		opTok := p.currTok
		p.nextToken()

		right, err := p.parseTerm()
		if err != nil {
			return nil, err
		}

		node := &rules.ExpressionGroup{}
		if opTok.Type == TKN_OR {
			node.Op = rules.OR
		} else {
			node.Op = rules.XOR
		}

		node.LeftExpressionGroup = left
		node.RightExpressionGroup = right

		left = node
	}

	return left, nil
}

func (p *Parser) parseTerm() (*rules.ExpressionGroup, error) {
	left, err := p.parseFactor()
	if err != nil {
		return nil, err
	}

	for p.currTok.Type == TKN_AND {
		p.nextToken()

		right, err := p.parseFactor()
		if err != nil {
			return nil, err
		}
		node := &rules.ExpressionGroup{
			Op:                   rules.AND,
			LeftExpressionGroup:  left,
			RightExpressionGroup: right,
		}
		left = node
	}

	return left, nil
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
		node := &rules.ExpressionGroup{
			Op: rules.NOTHING,
			LeftVariable: &rules.Variable{
				Letter: rune(p.currTok.Value[0]),
				Not:    false,
			},
		}
		p.nextToken()
		return node, nil
	}

	return nil, fmt.Errorf("expected variable, '(', '!' but got %v", p.currTok)
}

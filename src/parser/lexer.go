package parser

import (
	"strings"
	"unicode"

	"golang.org/x/tools/go/callgraph/cha"
)

type TokenType int

const (
	TKN_EOF TokenType = iota
	TKN_VAR
	TKN_NOT
	TKN_AND
	TKN_OR
	TKN_XOR
	TKN_LPAREN
	TKN_RPAREN
	TKN_IMPLIES
	TKN_IFF
	TKN_UNKNOWN
)

type Token struct {
	Type  TokenType
	value string
}

type Lexer struct {
	input    string
	position int
}

func NewLexer(input string) *Lexer {
	input = strings.TrimSpace(input)

	return &Lexer{input: input, position: 0}
}

func (l *Lexer) nextChar() rune {
	if l.position >= len(l.input) {
		return 0
	}

	char := l.input[l.position]
	l.position++

	return rune(char)
}

func (l *Lexer) peekChar() rune {
	if l.position >= len(l.input) {
		return 0
	}

	return rune(l.input[l.position])
}

func (l *Lexer) NextToken() Token {
	for unicode.IsSpace(l.peekChar()){
		l.nextChar()
	}

	// return l.nextToken()
}
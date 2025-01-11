package parser

import (
	"strings"
	"unicode"
)

type TokenType int

const (
	TKN_EOF     TokenType = iota
	TKN_VAR               // A..Z
	TKN_NOT               // '!'
	TKN_AND               // '+'
	TKN_OR                // '|'
	TKN_XOR               // '^'
	TKN_LPAREN            // '('
	TKN_RPAREN            // ')'
	TKN_IMPLIES           // '=>'
	TKN_IFF               // '<=>'
	TKN_UNKNOWN
)

type Token struct {
	Type  TokenType
	Value string
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
	for unicode.IsSpace(l.peekChar()) {
		l.nextChar()
	}

	char := l.nextChar()

	switch char {
	case 0:
		return Token{Type: TKN_EOF, Value: ""}
	case '+':
		return Token{Type: TKN_AND, Value: "+"}
	case '|':
		return Token{Type: TKN_OR, Value: "|"}
	case '^':
		return Token{Type: TKN_XOR, Value: "^"}
	case '!':
		return Token{Type: TKN_NOT, Value: "!"}
	case '(':
		return Token{Type: TKN_LPAREN, Value: "("}
	case ')':
		return Token{Type: TKN_RPAREN, Value: ")"}
	case '<':
		if l.peekChar() == '=' {
			l.nextChar()
			if l.peekChar() == '>' {
				l.nextChar()
				return Token{Type: TKN_IFF, Value: "<=>"}
			}
			return Token{Type: TKN_UNKNOWN, Value: "<"}
		}
		return Token{Type: TKN_UNKNOWN, Value: "<"}
	case '=':
		if l.peekChar() == '>' {
			l.nextChar()
			return Token{Type: TKN_IMPLIES, Value: "=>"}
		}
		return Token{Type: TKN_UNKNOWN, Value: "="}

	default:
		if unicode.IsLetter(char) {
			return Token{Type: TKN_VAR, Value: string(char)}
		}
		return Token{Type: TKN_UNKNOWN, Value: string(char)}
	}
}

package parser

import (
	"testing"
)

func TestLexerSimple(t *testing.T) {
	input := "A + !B => (C ^ (D + E))"
	lex := NewLexer(input)
	for {
		tok := lex.NextToken()
		t.Logf("Token: %v", tok)
		if tok.Type == TKN_EOF {
			break
		}
	}
}

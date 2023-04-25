package lexer

import (
	"testing"
)

func TestInc(t *testing.T) {
	lexer := NewLexer("ab", func(l *Lexer) LexingFunction { return nil })

	lexer.Inc()

	if lexer.Position != 1 {
		t.Errorf("")
	}

}

func TestDec(t *testing.T) {

}

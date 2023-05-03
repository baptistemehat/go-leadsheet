package lex

import (
	"testing"

	"github.com/baptistemehat/go-leadsheet/core/song/lexer/lexertoken"
	"github.com/stretchr/testify/assert"
)

func TestInc(t *testing.T) {

	input := "a"

	expectedToken := lexertoken.Token{
		Type:  lexertoken.TOKEN_EOF,
		Value: input,
	}

	lexer := NewLexer(input, nil)

	lexer.Inc()

	assert.Equal(t, uint(1), lexer.Position, "should be equal")
	assert.Equal(t, expectedToken, <-lexer.Tokens, "should be equal")
}

func TestInc_Blank(t *testing.T) {

	input := ""

	lexer := NewLexer(input, nil)

	lexer.Inc()

	assert.Equal(t, uint(1), lexer.Position, "should be equal")
	assert.Equal(t, lexertoken.TOKEN_ERROR, (<-lexer.Tokens).Type, "should be equal")
}
func TestDec(t *testing.T) {

}

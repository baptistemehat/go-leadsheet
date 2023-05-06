package lex

import (
	"testing"
	"unicode/utf8"

	"github.com/baptistemehat/go-leadsheet/core/song/lexer/lexertoken"
	"github.com/stretchr/testify/assert"
)

func TestInc_EOF(t *testing.T) {

	input := "a"

	expectedToken := lexertoken.Token{
		Type:  lexertoken.TOKEN_EOF,
		Value: input,
	}

	lexer := NewLexer(input, nil)

	lexer.Inc()

	assert.Equal(t, uint(1), lexer.Position, "should be equal")

	select {
	case actualToken := <-lexer.Tokens:
		assert.Equal(t, expectedToken, actualToken, "should be equal")
	default:
		t.Error("expected token from Tokens channel, got none")
	}

}

func TestInc_Overflow(t *testing.T) {

	lexer := NewLexer("", nil)

	lexer.Inc()

	assert.Equal(t, uint(1), lexer.Position, "should be equal")
	assert.Equal(t, lexertoken.TOKEN_ERROR, (<-lexer.Tokens).Type, "should be equal")
}

func TestDec(t *testing.T) {
	lexer := NewLexer("abcd", nil)
	lexer.Position = 2

	lexer.Dec()

	assert.Equal(t, uint(1), lexer.Position, "should be equal")
}

func TestDec_Underflow(t *testing.T) {
	lexer := NewLexer("", nil)

	lexer.Dec()

	assert.Equal(t, uint(0), lexer.Position, "should be equal")
}

func TestNextRune(t *testing.T) {

	input := "abcd"

	lexer := NewLexer(input, nil)

	actualRune := lexer.NextRune()

	assert.Equal(t, rune(input[0]), actualRune, "should be equal")
	assert.Equal(t, 1, lexer.Position, "should be equal")
}

func TestNextRune_LongRune(t *testing.T) {

	input := "äbcd"

	expectedRune, _ := utf8.DecodeRuneInString(input)

	lexer := NewLexer(input, nil)

	actualRune := lexer.NextRune()

	assert.Equal(t, expectedRune, actualRune, "should be equal")
	assert.Equal(t, 2, lexer.Position, "should be equal")
}

func TestNexRune_EOF(t *testing.T) {

	input := ""

	lexer := NewLexer(input, nil)

	actualRune := lexer.NextRune()

	assert.Equal(t, rune(0), actualRune, "should be equal")
	assert.Equal(t, 2, lexer.Position, "should be equal")
}

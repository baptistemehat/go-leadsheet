package lexingFunctions

import (
	"testing"
	"unicode/utf8"

	"github.com/baptistemehat/go-leadsheet/core/song/lexing"
	"github.com/stretchr/testify/assert"
)

var lexChordTestCases = []struct{
	name string
	input string
	expectedToken *lexing.Token
}{
	{
		name: "basic case - lex chord",
		input: "abc]",
		expectedToken: &lexing.Token{
			Type: lexing.TOKEN_CHORD,
			Value: "abc",
			Start: lexing.TokenPosition{Line: 0, Column: 0},
			End: lexing.TokenPosition{Line: 0, Column: 3},
		},
	},
	{
		name: "EOF",
		input: "abc",
		expectedToken: &lexing.Token{
			Type: lexing.TOKEN_ERROR,
			Value: lexing.LEXER_ERROR_UNEXPECTED_EOF,
			Start: lexing.TokenPosition{Line: 0, Column: 0},
			End: lexing.TokenPosition{Line: 0, Column: 3},
		},
	},
	{
		name: "rune error",
		input: "abc" + string(utf8.RuneError),
		expectedToken: &lexing.Token{
			Type: lexing.TOKEN_ERROR,
			Value: lexing.LEXER_ERROR_UNEXPECTED_RUNE,
			Start: lexing.TokenPosition{Line: 0, Column: 0},
			End: lexing.TokenPosition{Line: 0, Column: 3},
		},
	},
	{
		name: "newline error",
		input: "ab\nc]",
		expectedToken: &lexing.Token{
			Type: lexing.TOKEN_ERROR,
			Value: lexing.LEXER_ERROR_UNEXPECTED_NEWLINE,
			Start: lexing.TokenPosition{Line: 0, Column: 0},
			End: lexing.TokenPosition{Line: 0, Column: 2},
		},
	},
}
func TestLexChord(t *testing.T) {
	for _, testCase := range lexChordTestCases {
		t.Run(testCase.name, func(t *testing.T) {
			lexer := lexing.NewLexer(testCase.input, LexChord)

			actualToken := lexer.NextToken()

			assert.Equal(t, *testCase.expectedToken, actualToken)
		})
	}
}
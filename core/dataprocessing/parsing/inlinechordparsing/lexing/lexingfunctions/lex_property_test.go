package lexingfunctions

import (
	"testing"
	"unicode/utf8"

	"github.com/baptistemehat/go-leadsheet/core/dataprocessing/parsing/inlinechordparsing/lexing"
	"github.com/stretchr/testify/assert"
)

var lexPropertyKeyTestCases = []struct {
	name          string
	input         string
	expectedToken *lexing.Token
}{
	{
		name:  "basic case - lex property key",
		input: "abc:",
		expectedToken: &lexing.Token{
			Type:  lexing.TOKEN_PROPERTY_KEY,
			Value: "abc",
			Start: lexing.TokenPosition{Line: 0, Column: 0},
			End:   lexing.TokenPosition{Line: 0, Column: 3},
		},
	},
	{
		name:  "EOF",
		input: "abc",
		expectedToken: &lexing.Token{
			Type:  lexing.TOKEN_ERROR,
			Value: lexing.LEXER_ERROR_UNEXPECTED_EOF,
			Start: lexing.TokenPosition{Line: 0, Column: 0},
			End:   lexing.TokenPosition{Line: 0, Column: 3},
		},
	},
	{
		name:  "rune error",
		input: "abc" + string(utf8.RuneError),
		expectedToken: &lexing.Token{
			Type:  lexing.TOKEN_ERROR,
			Value: lexing.LEXER_ERROR_UNEXPECTED_RUNE,
			Start: lexing.TokenPosition{Line: 0, Column: 0},
			End:   lexing.TokenPosition{Line: 0, Column: 3},
		},
	},
	{
		name:  "newline error",
		input: "ab\nc:",
		expectedToken: &lexing.Token{
			Type:  lexing.TOKEN_ERROR,
			Value: lexing.LEXER_ERROR_UNEXPECTED_NEWLINE,
			Start: lexing.TokenPosition{Line: 0, Column: 0},
			End:   lexing.TokenPosition{Line: 0, Column: 2},
		},
	},
}

func TestLexPropertyKey(t *testing.T) {
	for _, testCase := range lexPropertyKeyTestCases {
		t.Run(testCase.name, func(t *testing.T) {
			lexer := lexing.NewLexer(testCase.input, LexPropertyKey)

			actualToken := lexer.NextToken()

			assert.Equal(t, *testCase.expectedToken, actualToken)
		})
	}
}

var lexPropertyValueTestCases = []struct {
	name          string
	input         string
	expectedToken *lexing.Token
}{
	{
		name:  "basic case - lex property value",
		input: "abc\ndef",
		expectedToken: &lexing.Token{
			Type:  lexing.TOKEN_PROPERTY_VALUE,
			Value: "abc",
			Start: lexing.TokenPosition{Line: 0, Column: 0},
			End:   lexing.TokenPosition{Line: 0, Column: 3},
		},
	},
	{
		name:  "EOF",
		input: "abc",
		expectedToken: &lexing.Token{
			Type:  lexing.TOKEN_ERROR,
			Value: lexing.LEXER_ERROR_UNEXPECTED_EOF,
			Start: lexing.TokenPosition{Line: 0, Column: 0},
			End:   lexing.TokenPosition{Line: 0, Column: 3},
		},
	},
	{
		name:  "rune error",
		input: "abc" + string(utf8.RuneError),
		expectedToken: &lexing.Token{
			Type:  lexing.TOKEN_ERROR,
			Value: lexing.LEXER_ERROR_UNEXPECTED_RUNE,
			Start: lexing.TokenPosition{Line: 0, Column: 0},
			End:   lexing.TokenPosition{Line: 0, Column: 3},
		},
	},
}

func TestLexPropertyValue(t *testing.T) {
	for _, testCase := range lexPropertyValueTestCases {
		t.Run(testCase.name, func(t *testing.T) {
			lexer := lexing.NewLexer(testCase.input, LexPropertyValue)

			actualToken := lexer.NextToken()

			assert.Equal(t, *testCase.expectedToken, actualToken)
		})
	}
}

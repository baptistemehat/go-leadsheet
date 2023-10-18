package lexingfunctions

import (
	"testing"
	"unicode/utf8"

	"github.com/baptistemehat/go-leadsheet/core/dataprocessing/parsing/inlinechordparsing/lexing"
	"github.com/stretchr/testify/assert"
)

var lexSectionNameTestCases = []struct {
	name          string
	input         string
	expectedToken *lexing.Token
}{
	{
		name:  "basic case - brace",
		input: "abc}",
		expectedToken: &lexing.Token{
			Type:  lexing.TOKEN_SECTION_NAME,
			Value: "abc",
			Start: lexing.TokenPosition{Line: 0, Column: 0},
			End:   lexing.TokenPosition{Line: 0, Column: 3},
		},
	},
	{
		name:  "EOF",
		input: "",
		expectedToken: &lexing.Token{
			Type:  lexing.TOKEN_ERROR,
			Value: lexing.LEXER_ERROR_UNEXPECTED_EOF,
			Start: lexing.TokenPosition{Line: 0, Column: 0},
			End:   lexing.TokenPosition{Line: 0, Column: 0},
		},
	},
	{
		name:  "rune error",
		input: string(utf8.RuneError),
		expectedToken: &lexing.Token{
			Type:  lexing.TOKEN_ERROR,
			Value: lexing.LEXER_ERROR_UNEXPECTED_RUNE,
			Start: lexing.TokenPosition{Line: 0, Column: 0},
			End:   lexing.TokenPosition{Line: 0, Column: 0},
		},
	},
}

func TestLexSectionName(t *testing.T) {
	for _, testCase := range lexSectionNameTestCases {
		t.Run(testCase.name, func(t *testing.T) {
			lexer := lexing.NewLexer(testCase.input, LexSectionName)

			actualToken := lexer.NextToken()

			assert.Equal(t, *testCase.expectedToken, actualToken)
		})
	}

}

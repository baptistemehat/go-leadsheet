package lexingfunctions

import (
	"testing"
	"unicode/utf8"

	"github.com/baptistemehat/go-leadsheet/core/dataprocessing/parsing/inlinechordparsing/lexing"
	"github.com/stretchr/testify/assert"
)

var lexRootTestCases = []struct {
	name          string
	input         string
	expectedToken *lexing.Token
}{
	// this test cannot be easily tested since EOF is alredy detected by SkipZhitespace
	// maybe this case can be removed from the function
	// 	{
	// 		name:  "EOF",
	// 		input: "\n",
	// 		expectedToken: &lexing.Token{
	// 			Type:  lexing.TOKEN_ERROR,
	// 			Value: lexing.LEXER_ERROR_UNEXPECTED_EOF,
	// 			Start: lexing.TokenPosition{Line: 1, Column: 0},
	// 			End:   lexing.TokenPosition{Line: 1, Column: 0},
	// 		},
	// 	},
	{
		name:  "rune error",
		input: "\n" + string(utf8.RuneError),
		expectedToken: &lexing.Token{
			Type:  lexing.TOKEN_ERROR,
			Value: lexing.LEXER_ERROR_UNEXPECTED_RUNE,
			Start: lexing.TokenPosition{Line: 1, Column: 0},
			End:   lexing.TokenPosition{Line: 1, Column: 0},
		},
	},
}

func TestLexRoot(t *testing.T) {
	for _, testCase := range lexRootTestCases {
		t.Run(testCase.name, func(t *testing.T) {
			// setup
			lexer := lexing.NewLexer(testCase.input, LexPropertyValue)
			lexer.NextToken()

			actualToken := lexer.NextToken()

			assert.Equal(t, *testCase.expectedToken, actualToken)
		})
	}

}

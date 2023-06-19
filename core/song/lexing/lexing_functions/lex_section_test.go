package lexingFunctions

import (
	"testing"
	"unicode/utf8"

	"github.com/baptistemehat/go-leadsheet/core/song/lexing"
	"github.com/stretchr/testify/assert"
)

// TODO this way of testing, makes the testing difficult, since using NextToken
// will call subsequent LexingFunctions without pushing a token

// TODO : for all LexingFunctions, we should be able to test the output, aka the returned lexing functinons
// for now, these tests are functionnal tests (which is fine, but incomplete if we do not add some unit tests)

// function comparison is not possible in golang, but we can compare function pointer
// with reflect.
// an in the best case, we would refactor to avoid this testing issue

var lexSectionTestCases = []struct {
	name          string
	input         string
	expectedToken *lexing.Token
}{
	{
		name:  "EOF",
		input: "",
		expectedToken: &lexing.Token{
			Type:  lexing.TOKEN_EOF,
			Value: "",
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

func TestLexSection(t *testing.T) {
	for _, testCase := range lexSectionTestCases {
		t.Run(testCase.name, func(t *testing.T) {
			lexer := lexing.NewLexer(testCase.input, LexSection)

			actualToken := lexer.NextToken()

			assert.Equal(t, *testCase.expectedToken, actualToken)
		})
	}
}

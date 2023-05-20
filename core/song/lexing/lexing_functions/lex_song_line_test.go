package lexingFunctions

import (
	"testing"
	"unicode/utf8"

	"github.com/baptistemehat/go-leadsheet/core/song/lexing"
	"github.com/stretchr/testify/assert"
)

var lexSongLineTestCases = []struct {
	name           string
	input          string
	expectedTokens []*lexing.Token
}{
	{
		name:  "basic case - lex song line - newline",
		input: "abc\n",
		expectedTokens: []*lexing.Token{
			{
				Type:  lexing.TOKEN_LYRICS,
				Value: "abc",
				Start: lexing.TokenPosition{Line: 0, Column: 0},
				End:   lexing.TokenPosition{Line: 0, Column: 3},
			},
		},
	},

	{
		name:  "basic case - lex song line - EOF",
		input: "abc",
		expectedTokens: []*lexing.Token{
			{
				Type:  lexing.TOKEN_LYRICS,
				Value: "abc",
				Start: lexing.TokenPosition{Line: 0, Column: 0},
				End:   lexing.TokenPosition{Line: 0, Column: 3},
			},
			{
				Type:  lexing.TOKEN_EOF,
				Value: "",
				Start: lexing.TokenPosition{Line: 0, Column: 3},
				End:   lexing.TokenPosition{Line: 0, Column: 3},
			},
		},
	},
	{
		name:  "rune error",
		input: "abc" + string(utf8.RuneError),
		expectedTokens: []*lexing.Token{
			{
				Type:  lexing.TOKEN_ERROR,
				Value: lexing.LEXER_ERROR_UNEXPECTED_RUNE,
				Start: lexing.TokenPosition{Line: 0, Column: 0},
				End:   lexing.TokenPosition{Line: 0, Column: 3},
			},
		},
	},
	{
		name:  "lex song line - left bracket",
		input: "abc[",
		expectedTokens: []*lexing.Token{
			{
				Type:  lexing.TOKEN_LYRICS,
				Value: "abc",
				Start: lexing.TokenPosition{Line: 0, Column: 0},
				End:   lexing.TokenPosition{Line: 0, Column: 3},
			},
		},
	},
	{
		name:  "lex song line - left brace",
		input: "abc{",
		expectedTokens: []*lexing.Token{
			{
				Type:  lexing.TOKEN_ERROR,
				Value: lexing.LEXER_ERROR_MISSING_NEWLINE_BEFORE_LEFT_BRACE,
				Start: lexing.TokenPosition{Line: 0, Column: 0},
				End:   lexing.TokenPosition{Line: 0, Column: 3},
			},
		},
	},
	{
		name:  "lex song line - left parenthesis",
		input: "abc(",
		expectedTokens: []*lexing.Token{
			{
				Type:  lexing.TOKEN_LYRICS,
				Value: "abc",
				Start: lexing.TokenPosition{Line: 0, Column: 0},
				End:   lexing.TokenPosition{Line: 0, Column: 3},
			},
		},
	},
}

func TestLexSongLine(t *testing.T) {
	for _, testCase := range lexSongLineTestCases {
		t.Run(testCase.name, func(t *testing.T) {
			lexer := lexing.NewLexer(testCase.input, LexSongLine)

			for _, expectedToken := range testCase.expectedTokens {
				actualToken := lexer.NextToken()
				assert.Equal(t, *expectedToken, actualToken)
			}
		})
	}
}

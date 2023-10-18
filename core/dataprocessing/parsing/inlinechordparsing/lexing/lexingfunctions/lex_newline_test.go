package lexingfunctions

import (
	"testing"

	"github.com/baptistemehat/go-leadsheet/core/dataprocessing/parsing/inlinechordparsing/lexing"
	"github.com/stretchr/testify/assert"
)

func TestLexNewline(t *testing.T) {

	expectedToken := lexing.Token{
		Type:  lexing.TOKEN_NEWLINE,
		Value: string(lexing.RUNE_NEWLINE),
		Start: lexing.TokenPosition{Line: 0, Column: 0},
		End:   lexing.TokenPosition{Line: 0, Column: 1},
	}

	lexer := lexing.NewLexer(string(lexing.RUNE_NEWLINE), LexNewLine)

	actualToken := lexer.NextToken()

	assert.Equal(t, expectedToken, actualToken)
}

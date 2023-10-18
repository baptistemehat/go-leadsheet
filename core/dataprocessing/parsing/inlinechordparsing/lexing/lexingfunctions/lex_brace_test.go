package lexingfunctions

import (
	"testing"

	"github.com/baptistemehat/go-leadsheet/core/dataprocessing/parsing/inlinechordparsing/lexing"
	"github.com/stretchr/testify/assert"
)

func TestLexLeftBrace(t *testing.T) {

	expectedToken := lexing.Token{
		Type:  lexing.TOKEN_LEFT_BRACE,
		Value: string(lexing.RUNE_LEFT_BRACE),
		Start: lexing.TokenPosition{Line: 0, Column: 0},
		End:   lexing.TokenPosition{Line: 0, Column: 1},
	}

	lexer := lexing.NewLexer(string(lexing.RUNE_LEFT_BRACE), LexLeftBrace)

	actualToken := lexer.NextToken()

	assert.Equal(t, expectedToken, actualToken)
}

func TestLexRightBrace(t *testing.T) {

	expectedToken := lexing.Token{
		Type:  lexing.TOKEN_RIGHT_BRACE,
		Value: string(lexing.RUNE_RIGHT_BRACE),
		Start: lexing.TokenPosition{Line: 0, Column: 0},
		End:   lexing.TokenPosition{Line: 0, Column: 1},
	}

	lexer := lexing.NewLexer(string(lexing.RUNE_RIGHT_BRACE), LexRightBrace)

	actualToken := lexer.NextToken()

	assert.Equal(t, expectedToken, actualToken)
}

package lexingfunctions

import (
	"testing"

	"github.com/baptistemehat/go-leadsheet/core/dataprocessing/parsing/inlinechordparsing/lexing"
	"github.com/stretchr/testify/assert"
)

func TestLexLeftBracket(t *testing.T) {

	expectedToken := lexing.Token{
		Type:  lexing.TOKEN_LEFT_BRACKET,
		Value: string(lexing.RUNE_LEFT_BRACKET),
		Start: lexing.TokenPosition{Line: 0, Column: 0},
		End:   lexing.TokenPosition{Line: 0, Column: 1},
	}

	lexer := lexing.NewLexer(string(lexing.RUNE_LEFT_BRACKET), LexLeftBracket)

	actualToken := lexer.NextToken()

	assert.Equal(t, expectedToken, actualToken)
}

func TestLexRightBracket(t *testing.T) {

	expectedToken := lexing.Token{
		Type:  lexing.TOKEN_RIGHT_BRACKET,
		Value: string(lexing.RUNE_RIGHT_BRACKET),
		Start: lexing.TokenPosition{Line: 0, Column: 0},
		End:   lexing.TokenPosition{Line: 0, Column: 1},
	}

	lexer := lexing.NewLexer(string(lexing.RUNE_RIGHT_BRACKET), LexRightBracket)

	actualToken := lexer.NextToken()

	assert.Equal(t, expectedToken, actualToken)
}

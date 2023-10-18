package lexingfunctions

import (
	"testing"

	"github.com/baptistemehat/go-leadsheet/core/dataprocessing/parsing/inlinechordparsing/lexing"
	"github.com/stretchr/testify/assert"
)

func TestLexLeftParenthesis(t *testing.T) {

	expectedToken := lexing.Token{
		Type:  lexing.TOKEN_LEFT_PARENTHESIS,
		Value: string(lexing.RUNE_LEFT_PARENTHESIS),
		Start: lexing.TokenPosition{Line: 0, Column: 0},
		End:   lexing.TokenPosition{Line: 0, Column: 1},
	}

	lexer := lexing.NewLexer(string(lexing.RUNE_LEFT_PARENTHESIS), LexLeftParenthesis)

	actualToken := lexer.NextToken()

	assert.Equal(t, expectedToken, actualToken)
}

func TestLexRightParenthesis(t *testing.T) {

	expectedToken := lexing.Token{
		Type:  lexing.TOKEN_RIGHT_PARENTHESIS,
		Value: string(lexing.RUNE_RIGHT_PARENTHESIS),
		Start: lexing.TokenPosition{Line: 0, Column: 0},
		End:   lexing.TokenPosition{Line: 0, Column: 1},
	}

	lexer := lexing.NewLexer(string(lexing.RUNE_RIGHT_PARENTHESIS), LexRightParenthesis)

	actualToken := lexer.NextToken()

	assert.Equal(t, expectedToken, actualToken)
}

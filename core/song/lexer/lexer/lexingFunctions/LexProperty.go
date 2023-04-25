package lexingFunctions

import (
	"strings"

	"github.com/baptistemehat/go-leadsheet/core/song/lexer/lexer"
	"github.com/baptistemehat/go-leadsheet/core/song/lexer/lexererrors"
	"github.com/baptistemehat/go-leadsheet/core/song/lexer/lexertoken"
)

// LexPropertyKey lexes property keys
// Syntax: <PropertyKey>: <PropertyValue>
func LexPropertyKey(lexer *lexer.Lexer) lexer.LexingFunction {

	for {
		if strings.HasPrefix(lexer.Input[lexer.Position:], lexertoken.COLUMN) {
			lexer.PushToken(lexertoken.TOKEN_PROPERTY_KEY)
			return LexColumn
		}

		// TODO : only alphanumeric (ie, exclude \n)

		lexer.Inc()

		if lexer.IsEOF() {
			return lexer.Errorf(lexererrors.LEXER_ERROR_UNEXPECTED_EOF)
		}
	}
}

func LexPropertyValue(lexer *lexer.Lexer) lexer.LexingFunction {
	for {
		// a property value ends with a newline
		if strings.HasPrefix(lexer.Input[lexer.Position:], lexertoken.NEWLINE) {
			lexer.PushToken(lexertoken.TOKEN_PROPERTY_VALUE)
			return LexRoot
		}

		lexer.Inc()

		if lexer.IsEOF() {
			return lexer.Errorf(lexererrors.LEXER_ERROR_UNEXPECTED_EOF)

		}
	}
}

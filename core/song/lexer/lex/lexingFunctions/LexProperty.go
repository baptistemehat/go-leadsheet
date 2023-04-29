package lexingFunctions

import (
	"strings"

	"github.com/baptistemehat/go-leadsheet/core/song/lexer/lex"
	"github.com/baptistemehat/go-leadsheet/core/song/lexer/lexererrors"
	"github.com/baptistemehat/go-leadsheet/core/song/lexer/lexertoken"
)

// LexPropertyKey lexes property keys
// Syntax: <PropertyKey>: <PropertyValue>
func LexPropertyKey(lexer *lex.Lexer) lex.LexingFunction {

	for {
		// if next char is column ":" (ie. end of property key)
		if strings.HasPrefix(lexer.Input[lexer.Position:], lexertoken.COLUMN) {

			// push property key token
			lexer.PushToken(lexertoken.TOKEN_PROPERTY_KEY)

			// lex column
			return LexColumn
		}

		// TODO : only alphanumeric (ie, exclude \n, space); do this in all lexing functions

		// else increment position
		lexer.Inc()

		// if EOF, throw error
		if lexer.IsEOF() {
			return lexer.Errorf(lexererrors.LEXER_ERROR_UNEXPECTED_EOF)
		}
	}
}

// LexPropertyValue
func LexPropertyValue(lexer *lex.Lexer) lex.LexingFunction {
	for {
		// if new char i newline "\n" (ie. end of property value)
		if strings.HasPrefix(lexer.Input[lexer.Position:], lexertoken.NEWLINE) {

			// push property value token
			lexer.PushToken(lexertoken.TOKEN_PROPERTY_VALUE)

			// lex root (ie. property or section)
			return LexRoot
		}

		// else increment position
		lexer.Inc()

		// if EOF, throw error
		if lexer.IsEOF() {
			return lexer.Errorf(lexererrors.LEXER_ERROR_UNEXPECTED_EOF)
		}
	}
}

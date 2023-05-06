package lexingFunctions

import (
	"log"

	"github.com/baptistemehat/go-leadsheet/core/song/lexer/lex"
	"github.com/baptistemehat/go-leadsheet/core/song/lexer/lexertoken"
)

// LexPropertyKey lexes property keys
// Syntax: <PropertyKey>: <PropertyValue>
func LexPropertyKey(lexer *lex.Lexer) lex.LexingFunction {

	for {
		nextRune := lexer.PeekRune()

		log.Println("pk : " + string(nextRune))
		switch nextRune {

		case lexertoken.EOF:
			return lexer.Errorf("EOF")

		case lexertoken.ERROR:
			return lexer.Errorf("ERR")

		case lexertoken.COLUMN:
			lexer.PushToken(lexertoken.TOKEN_PROPERTY_KEY)
			return LexColumn
		}

		lexer.GoToNextRune(nextRune)
		// // if next char is column ":" (ie. end of property key)
		// if strings.HasPrefix(lexer.Input[lexer.Position:], lexertoken.COLUMN) {

		// 	// push property key token
		// 	lexer.PushToken(lexertoken.TOKEN_PROPERTY_KEY)

		// 	// lex column
		// 	return LexColumn
		// }

		// // TODO : only alphanumeric (ie, exclude \n, space); do this in all lexing functions

		// // else increment position
		// lexer.Inc()

		// // if EOF, throw error
		// if lexer.IsEOF() {
		// 	return lexer.Errorf(lexererrors.LEXER_ERROR_UNEXPECTED_EOF)
		// }
	}
}

// LexPropertyValue
func LexPropertyValue(lexer *lex.Lexer) lex.LexingFunction {
	for {
		nextRune := lexer.PeekRune()

		switch nextRune {

		case lexertoken.EOF:
			return lexer.Errorf("")

		case lexertoken.ERROR:
			return lexer.Errorf("")

		case lexertoken.NEWLINE:
			lexer.PushToken(lexertoken.TOKEN_PROPERTY_VALUE)
			return LexRoot
		}
		lexer.GoToNextRune(nextRune)

		// // if new char i newline "\n" (ie. end of property value)
		// if strings.HasPrefix(lexer.Input[lexer.Position:], lexertoken.NEWLINE) {

		// 	// push property value token
		// 	lexer.PushToken(lexertoken.TOKEN_PROPERTY_VALUE)

		// 	// lex root (ie. property or section)
		// 	return LexRoot
		// }

		// // else increment position
		// lexer.Inc()

		// // if EOF, throw error
		// if lexer.IsEOF() {
		// 	return lexer.Errorf(lexererrors.LEXER_ERROR_UNEXPECTED_EOF)
		// }
	}
}

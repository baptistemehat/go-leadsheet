package lexingFunctions

import (
	"github.com/baptistemehat/go-leadsheet/core/song/lexer/lex"
	"github.com/baptistemehat/go-leadsheet/core/song/lexer/lexertoken"
)

// LexPropertyKey lexes property keys
// Syntax: <PropertyKey>: <PropertyValue>
func LexPropertyKey(lexer *lex.Lexer) lex.LexingFunction {

	for {
		nextRune := lexer.PeekRune()

		switch nextRune {

		case lexertoken.EOF:
			// TODO : normalise error messages
			lexer.Errorf("unexpected EOF while parsing property key")
			return nil

		case lexertoken.ERROR:
			lexer.Errorf("unexpected character found")
			return nil

		case lexertoken.COLUMN:
			lexer.PushToken(lexertoken.TOKEN_PROPERTY_KEY)
			return LexColumn
		}

		lexer.MoveAfterRune(nextRune)
	}
}

// LexPropertyValue
func LexPropertyValue(lexer *lex.Lexer) lex.LexingFunction {
	for {
		nextRune := lexer.PeekRune()

		switch nextRune {

		case lexertoken.EOF:
			// TODO : normalise error messages
			lexer.Errorf("unexpected EOF while parsing property value")
			return nil

		case lexertoken.ERROR:
			lexer.Errorf("unexpected character found")
			return nil

		case lexertoken.NEWLINE:
			lexer.PushToken(lexertoken.TOKEN_PROPERTY_VALUE)
			return LexRoot
		}
		lexer.MoveAfterRune(nextRune)
	}
}

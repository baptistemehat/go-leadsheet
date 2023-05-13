package lexingFunctions

import (
	"github.com/baptistemehat/go-leadsheet/core/song/lexing"
)

// LexPropertyKey lexes property keys
// Syntax: <PropertyKey>: <PropertyValue>
func LexPropertyKey(lexer *lexing.Lexer) lexing.LexingFunction {

	for {
		nextRune := lexer.PeekRune()

		switch nextRune {

		case lexing.RUNE_EOF:
			// TODO : normalise error messages
			lexer.Errorf("unexpected EOF while parsing property key")
			return nil

		case lexing.RUNE_ERROR:
			lexer.Errorf("unexpected character found")
			return nil

		case lexing.RUNE_COLUMN:
			lexer.PushToken(lexing.TOKEN_PROPERTY_KEY)
			return LexColumn
		}

		lexer.MoveAfterRune(nextRune)
	}
}

// LexPropertyValue
func LexPropertyValue(lexer *lexing.Lexer) lexing.LexingFunction {
	for {
		nextRune := lexer.PeekRune()

		switch nextRune {

		case lexing.RUNE_EOF:
			// TODO : normalise error messages
			lexer.Errorf("unexpected EOF while parsing property value")
			return nil

		case lexing.RUNE_ERROR:
			lexer.Errorf("unexpected character found")
			return nil

		case lexing.RUNE_NEWLINE:
			lexer.PushToken(lexing.TOKEN_PROPERTY_VALUE)
			return LexRoot
		}
		lexer.MoveAfterRune(nextRune)
	}
}

package lexingfunctions

import (
	"github.com/baptistemehat/go-leadsheet/core/dataprocessing/parsing/inlinechordparsing/lexing"
)

// LexChord
func LexChord(lexer *lexing.Lexer) lexing.LexingFunction {
	for {
		nextRune := lexer.PeekRune()

		switch nextRune {

		case lexing.RUNE_EOF:
			lexer.Errorf(lexing.LEXER_ERROR_UNEXPECTED_EOF)
			return nil

		case lexing.RUNE_ERROR:
			lexer.Errorf(lexing.LEXER_ERROR_UNEXPECTED_RUNE)
			return nil

		case lexing.RUNE_NEWLINE:
			lexer.Errorf(lexing.LEXER_ERROR_UNEXPECTED_NEWLINE)
		case lexing.RUNE_RIGHT_BRACKET:
			lexer.PushToken(lexing.TOKEN_CHORD)
			return LexRightBracket
		}

		lexer.MoveAfterRune(nextRune)
	}
}

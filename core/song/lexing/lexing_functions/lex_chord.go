package lexingFunctions

import (
	"github.com/baptistemehat/go-leadsheet/core/song/lexing"
)

// LexChord
func LexChord(lexer *lexing.Lexer) lexing.LexingFunction {
	for {
		nextRune := lexer.PeekRune()

		switch nextRune {

		case lexing.RUNE_EOF:
			// TODO : normalise error messages
			lexer.Errorf("unexpected EOF while parsing Chord")
			return nil

		case lexing.RUNE_ERROR:
			lexer.Errorf("unexpected character found")
			return nil

		case lexing.RUNE_RIGHT_BRACKET:
			lexer.PushToken(lexing.TOKEN_CHORD)
			return LexRightBracket
		}

		lexer.MoveAfterRune(nextRune)
	}
}

package lexingFunctions

import (
	"github.com/baptistemehat/go-leadsheet/core/song/lexer/lex"
	"github.com/baptistemehat/go-leadsheet/core/song/lexer/lexertoken"
)

// LexChord
func LexChord(lexer *lex.Lexer) lex.LexingFunction {
	for {
		nextRune := lexer.PeekRune()

		switch nextRune {

		case lexertoken.EOF:
			// TODO : normalise error messages
			lexer.Errorf("unexpected EOF while parsing Chord")
			return nil

		case lexertoken.ERROR:
			lexer.Errorf("unexpected character found")
			return nil

		case lexertoken.RIGHT_BRACKET:
			lexer.PushToken(lexertoken.TOKEN_CHORD)
			return LexRightBracket
		}

		lexer.GoToNextRune(nextRune)
	}
}

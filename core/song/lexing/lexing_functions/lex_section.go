package lexingFunctions

import (
	"github.com/baptistemehat/go-leadsheet/core/song/lexing"
)

// LexSection
func LexSection(lexer *lexing.Lexer) lexing.LexingFunction {

	nextRune := lexer.PeekRune()

	switch nextRune {
	case lexing.RUNE_EOF:
		lexer.PushToken(lexing.TOKEN_EOF)
		return nil

	case lexing.RUNE_ERROR:
		lexer.Errorf("unexpected character found")
		return nil

	case lexing.RUNE_LEFT_BRACE:
		return LexLeftBrace

	default:
		return LexSongLine
	}
}

package lexingFunctions

import (
	"github.com/baptistemehat/go-leadsheet/core/song/lexing"
)

// LexRoot
func LexRoot(lexer *lexing.Lexer) lexing.LexingFunction {
	// Skip whitespaces
	lexer.SkipWhitespace()

	nextRune := lexer.PeekRune()

	switch nextRune {

	case lexing.RUNE_EOF:
		// TODO : normalise error messages
		lexer.Errorf("unexpected EOF while parsing root ")
		return nil

	case lexing.RUNE_ERROR:
		lexer.Errorf("unexpected character found")
		return nil

	case lexing.RUNE_LEFT_BRACE:
		return LexLeftBrace

	default:
		return LexPropertyKey
	}
}

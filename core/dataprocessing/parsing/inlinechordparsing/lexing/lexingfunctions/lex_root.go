package lexingfunctions

import (
	"github.com/baptistemehat/go-leadsheet/core/dataprocessing/parsing/inlinechordparsing/lexing"
)

// LexRoot
func LexRoot(lexer *lexing.Lexer) lexing.LexingFunction {
	// Skip whitespaces
	lexer.SkipWhitespace()

	nextRune := lexer.PeekRune()

	switch nextRune {

	case lexing.RUNE_EOF:
		lexer.Errorf(lexing.LEXER_ERROR_UNEXPECTED_EOF)
		return nil

	case lexing.RUNE_ERROR:
		lexer.Errorf(lexing.LEXER_ERROR_UNEXPECTED_RUNE)
		return nil

	case lexing.RUNE_LEFT_BRACE:
		return LexLeftBrace

	default:
		return LexPropertyKey
	}
}

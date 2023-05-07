package lexingFunctions

import (
	"github.com/baptistemehat/go-leadsheet/core/song/lexer/lex"
	"github.com/baptistemehat/go-leadsheet/core/song/lexer/lexertoken"
)

// LexRoot
func LexRoot(lexer *lex.Lexer) lex.LexingFunction {
	// Skip whitespaces
	lexer.SkipWhitespace()

	nextRune := lexer.PeekRune()

	switch nextRune {

	case lexertoken.EOF:
		// TODO : normalise error messages
		lexer.Errorf("unexpected EOF while parsing root ")
		return nil

	case lexertoken.ERROR:
		lexer.Errorf("unexpected character found")
		return nil

	case lexertoken.LEFT_BRACE:
		return LexLeftBrace

	default:
		return LexPropertyKey
	}
}

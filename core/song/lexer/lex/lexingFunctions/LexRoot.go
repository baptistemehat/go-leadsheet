package lexingFunctions

import (
	"strings"

	"github.com/baptistemehat/go-leadsheet/core/song/lexer/lex"
	"github.com/baptistemehat/go-leadsheet/core/song/lexer/lexertoken"
)

// LexRoot
func LexRoot(lexer *lex.Lexer) lex.LexingFunction {
	// Skip whitespaces
	lexer.SkipWhitespace()

	// if next char is left brace "{" (ie. start of section)
	if strings.HasPrefix(lexer.Input[lexer.Position:], lexertoken.LEFT_BRACE) {

		// lex brace (ie. start of section)
		return LexLeftBrace

	} else {

		// else lex property key
		return LexPropertyKey
	}
}

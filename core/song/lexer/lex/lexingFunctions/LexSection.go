package lexingFunctions

import (
	"strings"

	"github.com/baptistemehat/go-leadsheet/core/song/lexer/lex"
	"github.com/baptistemehat/go-leadsheet/core/song/lexer/lexertoken"
)

// LexSection
func LexSection(lexer *lex.Lexer) lex.LexingFunction {

	// if next char is left brace "{" (ie.start of section)
	if strings.HasPrefix(lexer.Input[lexer.Position:], lexertoken.LEFT_BRACE) {

		// lex left brace (ie. start a new section)
		return LexLeftBrace

	} else {

		// else lex song line
		return LexSongLine
	}
}

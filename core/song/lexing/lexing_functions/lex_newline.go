package lexingFunctions

import (
	"github.com/baptistemehat/go-leadsheet/core/song/lexing"
)

// LexNewLine
func LexNewLine(lexer *lexing.Lexer) lexing.LexingFunction {
	lexer.MoveAfterRune(lexing.RUNE_NEWLINE)
	lexer.PushToken(lexing.TOKEN_NEWLINE)
	return LexSection
}

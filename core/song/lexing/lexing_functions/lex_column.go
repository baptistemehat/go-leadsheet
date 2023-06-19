package lexingFunctions

import (
	"github.com/baptistemehat/go-leadsheet/core/song/lexing"
)

// LexColumn
func LexColumn(lexer *lexing.Lexer) lexing.LexingFunction {
	lexer.MoveAfterRune(lexing.RUNE_COLUMN)
	lexer.PushToken(lexing.TOKEN_COLUMN)
	return LexPropertyValue
}

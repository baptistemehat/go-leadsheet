package lexingfunctions

import (
	"github.com/baptistemehat/go-leadsheet/core/dataprocessing/parsing/inlinechordparsing/lexing"
)

// LexColumn
func LexColumn(lexer *lexing.Lexer) lexing.LexingFunction {
	lexer.MoveAfterRune(lexing.RUNE_COLUMN)
	lexer.PushToken(lexing.TOKEN_COLUMN)
	return LexPropertyValue
}

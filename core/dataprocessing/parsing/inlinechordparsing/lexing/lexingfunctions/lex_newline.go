package lexingfunctions

import (
	"github.com/baptistemehat/go-leadsheet/core/dataprocessing/parsing/inlinechordparsing/lexing"
)

// LexNewLine
func LexNewLine(lexer *lexing.Lexer) lexing.LexingFunction {
	lexer.MoveAfterRune(lexing.RUNE_NEWLINE)
	lexer.PushToken(lexing.TOKEN_NEWLINE)
	return LexSection
}

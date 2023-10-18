package lexingfunctions

import (
	"github.com/baptistemehat/go-leadsheet/core/dataprocessing/parsing/inlinechordparsing/lexing"

)

// LexLeftBrace
func LexLeftBrace(lexer *lexing.Lexer) lexing.LexingFunction {
	lexer.MoveAfterRune(lexing.RUNE_LEFT_BRACE)
	lexer.PushToken(lexing.TOKEN_LEFT_BRACE)
	return LexSectionName
}

// LexRightBrace
func LexRightBrace(lexer *lexing.Lexer) lexing.LexingFunction {
	lexer.MoveAfterRune(lexing.RUNE_RIGHT_BRACE)
	lexer.PushToken(lexing.TOKEN_RIGHT_BRACE)
	return LexSection
}

package lexingFunctions

import (
	"github.com/baptistemehat/go-leadsheet/core/song/lexing"
)

// LexleftBracket
func LexLeftBracket(lexer *lexing.Lexer) lexing.LexingFunction {
	lexer.MoveAfterRune(lexing.RUNE_LEFT_BRACKET)
	lexer.PushToken(lexing.TOKEN_LEFT_BRACKET)
	return LexChord
}

// LexRightBracket
func LexRightBracket(lexer *lexing.Lexer) lexing.LexingFunction {
	lexer.MoveAfterRune(lexing.RUNE_RIGHT_BRACKET)
	lexer.PushToken(lexing.TOKEN_RIGHT_BRACKET)
	return LexSongLine
}

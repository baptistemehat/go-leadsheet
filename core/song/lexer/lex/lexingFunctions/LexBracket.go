package lexingFunctions

import (
	"github.com/baptistemehat/go-leadsheet/core/song/lexer/lex"
	"github.com/baptistemehat/go-leadsheet/core/song/lexer/lexertoken"
)

// LexleftBracket
func LexLeftBracket(lexer *lex.Lexer) lex.LexingFunction {
	lexer.MoveAfterRune(lexertoken.LEFT_BRACKET)
	lexer.PushToken(lexertoken.TOKEN_LEFT_BRACKET)
	return LexChord
}

// LexRightBracket
func LexRightBracket(lexer *lex.Lexer) lex.LexingFunction {
	lexer.MoveAfterRune(lexertoken.RIGHT_BRACKET)
	lexer.PushToken(lexertoken.TOKEN_RIGHT_BRACKET)
	return LexSongLine
}

package lexingFunctions

import (
	"github.com/baptistemehat/go-leadsheet/core/song/lexer/lexer"
	"github.com/baptistemehat/go-leadsheet/core/song/lexer/lexertoken"
)

func LexLeftBracket(lexer *lexer.Lexer) lexer.LexingFunction {
	lexer.Position += len(lexertoken.LEFT_BRACKET)
	lexer.PushToken(lexertoken.TOKEN_LEFT_BRACKET)
	return LexChord
}

func LexRightBracket(lexer *lexer.Lexer) lexer.LexingFunction {
	lexer.Position += len(lexertoken.RIGHT_BRACKET)
	lexer.PushToken(lexertoken.TOKEN_RIGHT_BRACKET)
	return LexSongLine
}

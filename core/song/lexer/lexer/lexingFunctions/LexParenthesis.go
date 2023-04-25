package lexingFunctions

import (
	"github.com/baptistemehat/go-leadsheet/core/song/lexer/lexer"
	"github.com/baptistemehat/go-leadsheet/core/song/lexer/lexertoken"
)

func LexLeftParenthesis(lexer *lexer.Lexer) lexer.LexingFunction {
	lexer.Position += len(lexertoken.LEFT_PARENTHESIS)
	lexer.PushToken(lexertoken.TOKEN_LEFT_PARENTHESIS)
	return LexWhisperedLyrics
}

func LexRightParenthesis(lexer *lexer.Lexer) lexer.LexingFunction {
	lexer.Position += len(lexertoken.RIGHT_PARENTHESIS)
	lexer.PushToken(lexertoken.TOKEN_RIGHT_PARENTHESIS)
	return LexSongLine
}

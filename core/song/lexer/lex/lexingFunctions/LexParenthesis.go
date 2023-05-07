package lexingFunctions

import (
	"github.com/baptistemehat/go-leadsheet/core/song/lexer/lex"
	"github.com/baptistemehat/go-leadsheet/core/song/lexer/lexertoken"
)

// LexLeftParenthesis
func LexLeftParenthesis(lexer *lex.Lexer) lex.LexingFunction {
	// TODO : maybe include GoToNextToken in PushToken, to be sure it is executed AFTER pushing token
	lexer.GoToNextRune(lexertoken.LEFT_PARENTHESIS)
	lexer.PushToken(lexertoken.TOKEN_LEFT_PARENTHESIS)
	return LexWhisperedLyrics
}

// LexRightParenthesis
func LexRightParenthesis(lexer *lex.Lexer) lex.LexingFunction {
	lexer.GoToNextRune(lexertoken.RIGHT_PARENTHESIS)
	lexer.PushToken(lexertoken.TOKEN_RIGHT_PARENTHESIS)
	return LexSongLine
}

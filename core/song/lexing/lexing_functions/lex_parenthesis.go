package lexingFunctions

import (
	"github.com/baptistemehat/go-leadsheet/core/song/lexing"
)

// LexLeftParenthesis
func LexLeftParenthesis(lexer *lexing.Lexer) lexing.LexingFunction {
	// TODO : maybe include GoToNextToken in PushToken, to be sure it is executed AFTER pushing token
	lexer.MoveAfterRune(lexing.RUNE_LEFT_PARENTHESIS)
	lexer.PushToken(lexing.TOKEN_LEFT_PARENTHESIS)
	return LexWhisperedLyrics
}

// LexRightParenthesis
func LexRightParenthesis(lexer *lexing.Lexer) lexing.LexingFunction {
	lexer.MoveAfterRune(lexing.RUNE_RIGHT_PARENTHESIS)
	lexer.PushToken(lexing.TOKEN_RIGHT_PARENTHESIS)
	return LexSongLine
}

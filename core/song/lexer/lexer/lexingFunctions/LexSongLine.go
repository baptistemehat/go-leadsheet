package lexingFunctions

import (
	"strings"

	"github.com/baptistemehat/go-leadsheet/core/song/lexer/lexer"
	"github.com/baptistemehat/go-leadsheet/core/song/lexer/lexererrors"
	"github.com/baptistemehat/go-leadsheet/core/song/lexer/lexertoken"
)

func LexSongLine(lexer *lexer.Lexer) lexer.LexingFunction {
	for {

		if strings.HasPrefix(lexer.InputToEnd(), lexertoken.NEWLINE) {
			lexer.PushToken(lexertoken.TOKEN_LYRICS)
			return LexNewLine

		} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.LEFT_BRACKET) {
			lexer.PushToken(lexertoken.TOKEN_LYRICS)
			return LexLeftBracket

		} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.LEFT_PARENTHESIS) {
			lexer.PushToken(lexertoken.TOKEN_LYRICS)
			return LexLeftParenthesis

		} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.LEFT_BRACE) {
			return lexer.Errorf(lexererrors.LEXER_ERROR_MISSING_NEWLINE_BEFORE_LEFT_BRACE)
		}

		lexer.Inc()

		// no bracket or brace or parenthesis

		if lexer.IsEOF() {
			return lexer.Errorf(lexererrors.LEXER_ERROR_MISSING_NEWLINE_BEFORE_EOF)
		}

	}
}

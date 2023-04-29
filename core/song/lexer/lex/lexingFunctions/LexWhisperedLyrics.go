package lexingFunctions

import (
	"strings"

	"github.com/baptistemehat/go-leadsheet/core/song/lexer/lex"
	"github.com/baptistemehat/go-leadsheet/core/song/lexer/lexererrors"
	"github.com/baptistemehat/go-leadsheet/core/song/lexer/lexertoken"
)

// TODO : create a dedicated token for whispered lyrics

// LexWhisperedLyrics
func LexWhisperedLyrics(lexer *lex.Lexer) lex.LexingFunction {
	for {
		// if next char is right parenthesis ")" (ie. end of whispered lyrics)
		if strings.HasPrefix(lexer.Input[lexer.Position:], lexertoken.RIGHT_PARENTHESIS) {

			// push lyrics token
			lexer.PushToken(lexertoken.TOKEN_LYRICS)
			// lex right parenthesis
			return LexRightParenthesis
		}

		// else increament position
		lexer.Inc()

		// if EOF, throw error
		if lexer.IsEOF() {
			return lexer.Errorf(lexererrors.LEXER_ERROR_MISSING_RIGHT_PARENTHESIS)
		}
	}
}

package lexingFunctions

import (
	"github.com/baptistemehat/go-leadsheet/core/song/lexer/lex"
	"github.com/baptistemehat/go-leadsheet/core/song/lexer/lexertoken"
)

// TODO : create a dedicated token for whispered lyrics

// LexWhisperedLyrics
func LexWhisperedLyrics(lexer *lex.Lexer) lex.LexingFunction {
	for {
		nextRune := lexer.PeekRune()

		switch nextRune {

		case lexertoken.EOF:
			return lexer.Errorf("")

		case lexertoken.ERROR:
			return lexer.Errorf("")

		case lexertoken.RIGHT_PARENTHESIS:
			lexer.PushToken(lexertoken.TOKEN_LYRICS)
			return LexRightParenthesis
		}

		lexer.GoToNextRune(nextRune)
		// // if next char is right parenthesis ")" (ie. end of whispered lyrics)
		// if strings.HasPrefix(lexer.Input[lexer.Position:], lexertoken.RIGHT_PARENTHESIS) {

		// 	// push lyrics token
		// 	lexer.PushToken(lexertoken.TOKEN_LYRICS)
		// 	// lex right parenthesis
		// 	return LexRightParenthesis
		// }

		// // else increament position
		// lexer.Inc()

		// // if EOF, throw error
		// if lexer.IsEOF() {
		// 	return lexer.Errorf(lexererrors.LEXER_ERROR_MISSING_RIGHT_PARENTHESIS)
		// }
	}
}

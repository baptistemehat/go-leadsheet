package lexingFunctions

import (
	"github.com/baptistemehat/go-leadsheet/core/song/lexer/lex"
	"github.com/baptistemehat/go-leadsheet/core/song/lexer/lexertoken"
)

// LexSectionName
func LexSectionName(lexer *lex.Lexer) lex.LexingFunction {
	for {

		// TODO : exclude non-alphanum char

		nextRune := lexer.PeekRune()

		switch nextRune {

		case lexertoken.EOF:
			return lexer.Errorf("")

		case lexertoken.ERROR:
			return lexer.Errorf("")

		case lexertoken.RIGHT_BRACE:
			lexer.PushToken(lexertoken.TOKEN_SECTION_NAME)
			return LexRightBrace
		}

		// TODO : rename
		lexer.GoToNextRune(nextRune)

		// // if next rune is right brace "}" (ie. end of section name)
		// if strings.HasPrefix(lexer.Input[lexer.Position:], lexertoken.RIGHT_BRACE) {

		// 	// push section name token
		// 	lexer.PushToken(lexertoken.TOKEN_SECTION_NAME)
		// 	// lex right brace
		// 	return LexRightBrace
		// }

		// // else increment position
		// lexer.Inc()

		// // if EOF, throw error
		// if lexer.IsEOF() {
		// 	return lexer.Errorf(lexererrors.LEXER_ERROR_MISSING_RIGHT_BRACE)
		// }
	}
}

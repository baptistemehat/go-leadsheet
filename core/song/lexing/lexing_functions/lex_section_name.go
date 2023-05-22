package lexingFunctions

import (
	"github.com/baptistemehat/go-leadsheet/core/song/lexing"
)

// LexSectionName
func LexSectionName(lexer *lexing.Lexer) lexing.LexingFunction {
	for {

		// TODO : exclude non-alphanum char

		nextRune := lexer.PeekRune()

		switch nextRune {

		case lexing.RUNE_EOF:
			lexer.Errorf(lexing.LEXER_ERROR_UNEXPECTED_EOF)
			return nil

		case lexing.RUNE_ERROR:
			lexer.Errorf(lexing.LEXER_ERROR_UNEXPECTED_RUNE)
			return nil

		case lexing.RUNE_RIGHT_BRACE:
			lexer.PushToken(lexing.TOKEN_SECTION_NAME)
			return LexRightBrace
		}

		lexer.MoveAfterRune(nextRune)
	}
}

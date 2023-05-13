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
			// TODO : normalise error messages
			lexer.Errorf("unexpected EOF while parsing section name: position")
			return nil

		case lexing.RUNE_ERROR:
			lexer.Errorf("unexpected character found")
			return nil

		case lexing.RUNE_RIGHT_BRACE:
			lexer.PushToken(lexing.TOKEN_SECTION_NAME)
			return LexRightBrace
		}

		// TODO : rename
		lexer.MoveAfterRune(nextRune)
	}
}

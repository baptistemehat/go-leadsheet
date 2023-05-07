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
			// TODO : normalise error messages
			lexer.Errorf("unexpected EOF while parsing section name: position")
			return nil

		case lexertoken.ERROR:
			lexer.Errorf("unexpected character found")
			return nil

		case lexertoken.RIGHT_BRACE:
			lexer.PushToken(lexertoken.TOKEN_SECTION_NAME)
			return LexRightBrace
		}

		// TODO : rename
		lexer.MoveAfterRune(nextRune)
	}
}

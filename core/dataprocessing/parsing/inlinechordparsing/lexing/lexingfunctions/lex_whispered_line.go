package lexingfunctions

import (
	"github.com/baptistemehat/go-leadsheet/core/dataprocessing/parsing/inlinechordparsing/lexing"
)

// TODO : create a dedicated token for whispered lyrics

// LexWhisperedLyrics
func LexWhisperedLyrics(lexer *lexing.Lexer) lexing.LexingFunction {
	for {
		nextRune := lexer.PeekRune()

		switch nextRune {

		case lexing.RUNE_EOF:
			// TODO : normalise error messages
			lexer.Errorf("unexpected EOF while parsing root")
			return nil

		case lexing.RUNE_ERROR:
			lexer.Errorf("unexpected character found")
			return nil

		case lexing.RUNE_RIGHT_PARENTHESIS:
			lexer.PushToken(lexing.TOKEN_LYRICS)
			return LexRightParenthesis
		}

		lexer.MoveAfterRune(nextRune)
	}
}

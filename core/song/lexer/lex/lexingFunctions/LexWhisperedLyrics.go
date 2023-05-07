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
			// TODO : normalise error messages
			lexer.Errorf("unexpected EOF while parsing root")
			return nil

		case lexertoken.ERROR:
			lexer.Errorf("unexpected character found")
			return nil

		case lexertoken.RIGHT_PARENTHESIS:
			lexer.PushToken(lexertoken.TOKEN_LYRICS)
			return LexRightParenthesis
		}

		lexer.MoveAfterRune(nextRune)
	}
}

package lexingFunctions

import (
	"github.com/baptistemehat/go-leadsheet/core/song/lexer/lex"
	"github.com/baptistemehat/go-leadsheet/core/song/lexer/lexertoken"
)

// LexSection
func LexSection(lexer *lex.Lexer) lex.LexingFunction {

	nextRune := lexer.PeekRune()

	switch nextRune {
	case lexertoken.EOF:
		lexer.PushToken(lexertoken.TOKEN_EOF)
		return nil

	case lexertoken.ERROR:
		lexer.Errorf("unexpected character found")
		return nil

	case lexertoken.LEFT_BRACE:
		return LexLeftBrace

	default:
		return LexSongLine
	}
}

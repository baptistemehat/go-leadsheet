package lexingFunctions

import (
	"log"

	"github.com/baptistemehat/go-leadsheet/core/song/lexer/lex"
	"github.com/baptistemehat/go-leadsheet/core/song/lexer/lexertoken"
)

// LexSection
func LexSection(lexer *lex.Lexer) lex.LexingFunction {

	nextRune := lexer.PeekRune()
	log.Println("LexSection : " + string(nextRune))

	switch nextRune {
	case lexertoken.EOF:
		lexer.PushToken(lexertoken.TOKEN_EOF)
		return nil

	case lexertoken.ERROR:
		return lexer.Errorf("")

	case lexertoken.LEFT_BRACE:
		return LexLeftBrace

	default:
		return LexSongLine
	}

	// // if next char is left brace "{" (ie.start of section)
	// if strings.HasPrefix(lexer.Input[lexer.Position:], lexertoken.LEFT_BRACE) {

	// 	// lex left brace (ie. start a new section)
	// 	return LexLeftBrace

	// } else {

	// 	// else lex song line
	// 	return LexSongLine
	// }
}

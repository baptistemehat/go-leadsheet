package lexingFunctions

import (
	"log"

	"github.com/baptistemehat/go-leadsheet/core/song/lexer/lex"
	"github.com/baptistemehat/go-leadsheet/core/song/lexer/lexererrors"
	"github.com/baptistemehat/go-leadsheet/core/song/lexer/lexertoken"
)

// LexRoot
func LexRoot(lexer *lex.Lexer) lex.LexingFunction {
	// Skip whitespaces
	lexer.SkipWhitespace()

	nextRune := lexer.PeekRune()

	log.Println("root nextRune : " + string(nextRune))
	switch nextRune {

	case lexertoken.EOF:
		return lexer.Errorf(lexererrors.LEXER_ERROR_UNEXPECTED_EOF)

	case lexertoken.ERROR:
		return lexer.Errorf("error while parsing rune")

	case lexertoken.LEFT_BRACE:
		return LexLeftBrace

	default:
		return LexPropertyKey
	}

	// // if next rune is left brace "{" (ie. start of section)
	// if lexer.NextRune() == lexertoken.LEFT_BRACE {

	// 	// lex brace (ie. start of section)
	// 	return LexLeftBrace

	// } else {

	// 	// else lex property key
	// 	return LexPropertyKey
	// }
}

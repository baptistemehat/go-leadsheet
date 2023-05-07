package lexertoken

import "fmt"

// TODO : add line and column to lexer

type TokenPosition struct {
	Line   int
	Column int
}

func NewToken() Token {
	return Token{
		Type:  TOKEN_UNKNOWN,
		Value: "",
		Start: TokenPosition{Line: 0, Column: 0},
		End:   TokenPosition{Line: 0, Column: 0},
	}
}
func (tokenPosition *TokenPosition) String() string {
	return fmt.Sprintf("[%d:%d]", tokenPosition.Line, tokenPosition.Column)
}

// Token
type Token struct {
	Type  TokenType
	Value string
	Start TokenPosition
	End   TokenPosition
}

// IsEOF
func (t Token) IsEOF() bool {
	return t.Type == TOKEN_EOF
}

// String
func (token Token) String() string {
	tokenValue := token.Value

	if token.Type == TOKEN_EOF {
		tokenValue = "EOF"
	}
	return fmt.Sprintf("start%s end%s %s %q ", token.Start.String(), token.End.String(), token.Type.String(), tokenValue)
}

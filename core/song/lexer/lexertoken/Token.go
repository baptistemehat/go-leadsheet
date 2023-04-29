package lexertoken

import "fmt"

// Token
type Token struct {
	Type  TokenType
	Value string
}

// IsEOF
func (t Token) IsEOF() bool {
	return t.Type == TOKEN_EOF
}

// String
func (t Token) String() string {
	switch t.Type {

	case TOKEN_EOF:
		return "EOF"

	case TOKEN_ERROR:
		return t.Value
	}

	return fmt.Sprintf("%q", t.Value)
}

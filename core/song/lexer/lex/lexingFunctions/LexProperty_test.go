package lexingFunctions

import (
	"testing"

	"github.com/baptistemehat/go-leadsheet/core/song/lexer/lex"
	"github.com/baptistemehat/go-leadsheet/core/song/lexer/lexertoken"
)

type testCase struct {
	Name          string
	Input         string
	LexingFunc    lex.LexingFunction
	ExpectedToken lexertoken.Token
}

var testCaseArray = []testCase{
	{
		Name:       "Key",
		Input:      "Title: Hotel California\n",
		LexingFunc: LexPropertyKey,
		ExpectedToken: lexertoken.Token{
			Type:  lexertoken.TOKEN_PROPERTY_KEY,
			Value: "Title",
		},
	},
	{
		Name:       "Key_EOF",
		Input:      "Title",
		LexingFunc: LexPropertyKey,
		ExpectedToken: lexertoken.Token{
			Type:  lexertoken.TOKEN_EOF,
			Value: "Title",
		},
	},
	{
		Name:       "Value",
		Input:      "Hotel California\n",
		LexingFunc: LexPropertyValue,
		ExpectedToken: lexertoken.Token{
			Type:  lexertoken.TOKEN_PROPERTY_VALUE,
			Value: "Hotel California",
		},
	},
	{
		Name:       "Value_EOF",
		Input:      "Hotel California",
		LexingFunc: LexPropertyValue,
		ExpectedToken: lexertoken.Token{
			Type:  lexertoken.TOKEN_EOF,
			Value: "Hotel California",
		},
	},
}

func TestLexProperty(t *testing.T) {

	for _, testCase := range testCaseArray {
		t.Run(testCase.Name, func(t *testing.T) {
			lexer := lex.NewLexer(testCase.Input, LexRoot)
			testCase.LexingFunc(lexer)

			//if actualToken := <-lexer.Tokens; actualToken != testCase.ExpectedToken {
			//	actualTokenJson, _ := json.Marshal(actualToken)
			//	expectedTokenJson, _ := json.Marshal(testCase.ExpectedToken)
			//	t.Errorf("Expected: %s, Got: %s", string(expectedTokenJson), string(actualTokenJson))
			//}
		})
	}
}

// TODO : this is a behavioural test, shoud not be here
// func TestLexPropertyFull(t *testing.T) {
// 	l := lexer.NewLexer(testInput)

// 	if token := l.NextToken(); token.Type != lexertoken.TOKEN_PROPERTY_KEY {
// 		t.Errorf("Property key")
// 	}

// 	if token := l.NextToken(); token.Type != lexertoken.TOKEN_COLUMN {
// 		t.Errorf("Column")
// 	}

// 	if token := l.NextToken(); token.Type != lexertoken.TOKEN_PROPERTY_VALUE {
// 		t.Errorf("Property value")
// 	}
// }

package lexing

import (
	"testing"
	"unicode/utf8"

	"github.com/stretchr/testify/assert"
)

// **********************
//     TEST TOKENS
// **********************

// testTokenType is a test token type
const testTokenType TokenType = -1

// testToken is a test token
var testToken = Token{
	Type:  testTokenType,
	Value: "test",
	Start: TokenPosition{Line: 0, Column: 0},
	End:   TokenPosition{Line: 0, Column: 12},
}

// **********************
//  LEXING FUNCTION MOCK
// **********************

// MockLexingFuncCallCounter
type MockLexingFuncCallCounter struct {
	callCount int
}

// ResetCount resets the call counter
func (mockLexingFuncCallCount *MockLexingFuncCallCounter) ResetCount() {
	mockLexingFuncCallCount.callCount = 0
}

// GetCount returns the call counter
func (mockLexingFuncCallCount *MockLexingFuncCallCounter) GetCount() int {
	return mockLexingFuncCallCount.callCount
}

// Increment increments the call counter
func (mockLexingFuncCallCounter *MockLexingFuncCallCounter) Increment() {
	mockLexingFuncCallCounter.callCount++
}

// mockLexingFuncCallCounter is the function call counter associated to
var mockLexingFuncCallCounter MockLexingFuncCallCounter

// mockLexingFunction mocks a lexing function.
// Increments the MockLexingFuncCallCounter and pushes a testToken in the Tokens channel.
func mockLexingFunction(lexer *Lexer) LexingFunction {
	lexer.tokens <- testToken
	mockLexingFuncCallCounter.Increment()
	return mockLexingFunction
}

// **********************
//     TEST FUNCTIONS
// **********************

// assertEqualLexer assert if two Lexers are equal.
// This function only asserts equality for fields implementing '=='.
// NextLexingFunction function and Tokens channel are not considred in this function.
func assertEqualLexer(t *testing.T, expectedLexer, actualLexer *Lexer) {
	assert.Equal(t, expectedLexer.input, actualLexer.input)
	assert.Equal(t, expectedLexer.currentToken, actualLexer.currentToken)
	assert.Equal(t, expectedLexer.currentTokenStart, actualLexer.currentTokenStart)
	assert.Equal(t, expectedLexer.position, actualLexer.position)
}

// **********************
//      TEST CASES
// **********************

var newLexerTestCase = struct {
	input          string
	lexingFunction LexingFunction
	expectedLexer  *Lexer
}{
	input:          "test input",
	lexingFunction: nil,
	expectedLexer: &Lexer{
		input:              "test input",
		tokens:             make(chan Token, 5),
		nextLexingFunction: nil,
		currentToken: Token{
			Type:  TOKEN_UNKNOWN,
			Value: "",
			Start: TokenPosition{Line: 0, Column: 0},
			End:   TokenPosition{Line: 0, Column: 0},
		},
		currentTokenStart: 0,
		position:          0,
	},
}

func TestNewLexer(t *testing.T) {
	actualLexer := NewLexer(newLexerTestCase.input, newLexerTestCase.lexingFunction)

	assertEqualLexer(t, newLexerTestCase.expectedLexer, actualLexer)
}

var moveAfterRunTestCases = []struct {
	name          string
	inputRune     rune
	inputLexer    *Lexer
	expectedLexer *Lexer
}{
	{
		name:      "one byte rune",
		inputRune: 'a',
		inputLexer: &Lexer{
			input:              "abcde",
			tokens:             make(chan Token, 5),
			nextLexingFunction: nil,
			currentToken: Token{
				Type:  TOKEN_UNKNOWN,
				Value: "",
				Start: TokenPosition{Line: 0, Column: 2},
				End:   TokenPosition{Line: 0, Column: 4},
			},
			currentTokenStart: 2,
			position:          4,
		},
		expectedLexer: &Lexer{
			input:              "abcde",
			tokens:             make(chan Token, 5),
			nextLexingFunction: nil,
			currentToken: Token{
				Type:  TOKEN_UNKNOWN,
				Value: "",
				Start: TokenPosition{Line: 0, Column: 2},
				End:   TokenPosition{Line: 0, Column: 5},
			},
			currentTokenStart: 2,
			position:          5,
		}},
	{
		name:      "two byte rune",
		inputRune: 'ä',
		inputLexer: &Lexer{
			input:              "äbcde",
			tokens:             make(chan Token, 5),
			nextLexingFunction: nil,
			currentToken: Token{
				Type:  TOKEN_UNKNOWN,
				Value: "",
				Start: TokenPosition{Line: 0, Column: 0},
				End:   TokenPosition{Line: 0, Column: 0},
			},
			currentTokenStart: 0,
			position:          0,
		},
		expectedLexer: &Lexer{
			input:              "äbcde",
			tokens:             make(chan Token, 5),
			nextLexingFunction: nil,
			currentToken: Token{
				Type:  TOKEN_UNKNOWN,
				Value: "",
				Start: TokenPosition{Line: 0, Column: 0},
				End:   TokenPosition{Line: 0, Column: 1},
			},
			currentTokenStart: 0,
			position:          2,
		}},
	{
		name:      "newline",
		inputRune: RUNE_NEWLINE,
		inputLexer: &Lexer{
			input:              "\n",
			tokens:             make(chan Token, 5),
			nextLexingFunction: nil,
			currentToken: Token{
				Type:  TOKEN_UNKNOWN,
				Value: "",
				Start: TokenPosition{Line: 0, Column: 0},
				End:   TokenPosition{Line: 0, Column: 0},
			},
			currentTokenStart: 0,
			position:          0,
		},
		expectedLexer: &Lexer{
			input:              "\n",
			tokens:             make(chan Token, 5),
			nextLexingFunction: nil,
			currentToken: Token{
				Type:  TOKEN_UNKNOWN,
				Value: "",
				Start: TokenPosition{Line: 1, Column: 0},
				End:   TokenPosition{Line: 1, Column: 0},
			},
			currentTokenStart: 0,
			position:          1,
		}},
	{
		name:      "EOF rune",
		inputRune: RUNE_EOF,
		inputLexer: &Lexer{
			input:              "",
			tokens:             make(chan Token, 5),
			nextLexingFunction: nil,
			currentToken: Token{
				Type:  TOKEN_UNKNOWN,
				Value: "",
				Start: TokenPosition{Line: 0, Column: 0},
				End:   TokenPosition{Line: 0, Column: 0},
			},
			currentTokenStart: 0,
			position:          0,
		},
		expectedLexer: &Lexer{
			input:              "",
			tokens:             make(chan Token, 5),
			nextLexingFunction: nil,
			currentToken: Token{
				Type:  TOKEN_UNKNOWN,
				Value: "",
				Start: TokenPosition{Line: 0, Column: 0},
				End:   TokenPosition{Line: 0, Column: 0},
			},
			currentTokenStart: 0,
			position:          0,
		}},
	{
		name:      "ERROR rune",
		inputRune: RUNE_ERROR,
		inputLexer: &Lexer{
			input:              "",
			tokens:             make(chan Token, 5),
			nextLexingFunction: nil,
			currentToken: Token{
				Type:  TOKEN_UNKNOWN,
				Value: "",
				Start: TokenPosition{Line: 0, Column: 0},
				End:   TokenPosition{Line: 0, Column: 0},
			},
			currentTokenStart: 0,
			position:          0,
		},
		expectedLexer: &Lexer{
			input:              "",
			tokens:             make(chan Token, 5),
			nextLexingFunction: nil,
			currentToken: Token{
				Type:  TOKEN_UNKNOWN,
				Value: "",
				Start: TokenPosition{Line: 0, Column: 0},
				End:   TokenPosition{Line: 0, Column: 0},
			},
			currentTokenStart: 0,
			position:          0,
		}},
}

func TestMoveAfterRune(t *testing.T) {
	for _, testCase := range moveAfterRunTestCases {
		t.Run(testCase.name, func(t *testing.T) {

			testCase.inputLexer.MoveAfterRune(testCase.inputRune)

			assertEqualLexer(t, testCase.expectedLexer, testCase.inputLexer)

		})
	}
}

var peekRuneTestCases = []struct {
	name          string
	expectedRune  rune
	expectedLexer *Lexer
}{
	{
		name:         "one byte rune",
		expectedRune: 'c',
		expectedLexer: &Lexer{
			input:              "abcde",
			tokens:             make(chan Token, 5),
			nextLexingFunction: nil,
			currentToken: Token{
				Type:  TOKEN_UNKNOWN,
				Value: "ab",
				Start: TokenPosition{Line: 0, Column: 0},
				End:   TokenPosition{Line: 0, Column: 2},
			},
			currentTokenStart: 0,
			position:          2,
		}},
	{
		name:         "two byte rune",
		expectedRune: 'ä',
		expectedLexer: &Lexer{
			input:              "äbcde",
			tokens:             make(chan Token, 5),
			nextLexingFunction: nil,
			currentToken: Token{
				Type:  TOKEN_UNKNOWN,
				Value: "",
				Start: TokenPosition{Line: 0, Column: 0},
				End:   TokenPosition{Line: 0, Column: 0},
			},
			currentTokenStart: 0,
			position:          0,
		}},
	{
		name:         "empty input",
		expectedRune: RUNE_EOF,
		expectedLexer: &Lexer{
			input:              "",
			tokens:             make(chan Token, 5),
			nextLexingFunction: nil,
			currentToken: Token{
				Type:  TOKEN_UNKNOWN,
				Value: "",
				Start: TokenPosition{Line: 0, Column: 0},
				End:   TokenPosition{Line: 0, Column: 0},
			},
			currentTokenStart: 0,
			position:          0,
		}},
	{
		name:         "EOF",
		expectedRune: RUNE_EOF,
		expectedLexer: &Lexer{
			input:              "abcde",
			tokens:             make(chan Token, 5),
			nextLexingFunction: nil,
			currentToken: Token{
				Type:  TOKEN_UNKNOWN,
				Value: "ab",
				Start: TokenPosition{Line: 0, Column: 0},
				End:   TokenPosition{Line: 0, Column: 5},
			},
			currentTokenStart: 0,
			position:          5,
		}},
	{
		name:         "RuneError",
		expectedRune: RUNE_ERROR,
		expectedLexer: &Lexer{
			input:              string(utf8.RuneError),
			tokens:             make(chan Token, 5),
			nextLexingFunction: nil,
			currentToken: Token{
				Type:  TOKEN_UNKNOWN,
				Value: "",
				Start: TokenPosition{Line: 0, Column: 0},
				End:   TokenPosition{Line: 0, Column: 0},
			},
			currentTokenStart: 0,
			position:          0,
		}},
}

func TestPeekRune(t *testing.T) {
	for _, testCase := range peekRuneTestCases {
		t.Run(testCase.name, func(t *testing.T) {
			lexer := *testCase.expectedLexer

			actualRune := lexer.PeekRune()

			assert.Equal(t, testCase.expectedRune, actualRune, "should be equal")
			assertEqualLexer(t, testCase.expectedLexer, &lexer)
		})
	}
}

var pushTokenTestCases = []struct {
	name           string
	inputTokenType TokenType
	expectedToken  Token
	inputLexer     *Lexer
	expectedLexer  *Lexer
}{
	{
		name:           "basic case - push lyrics token",
		inputTokenType: TOKEN_LYRICS,
		expectedToken: Token{
			Type:  TOKEN_LYRICS,
			Value: "bcd",
			Start: TokenPosition{Line: 0, Column: 1},
			End:   TokenPosition{Line: 0, Column: 4},
		},
		inputLexer: &Lexer{
			input:              "abcde",
			tokens:             make(chan Token, 5),
			nextLexingFunction: nil,
			currentToken: Token{
				Type:  TOKEN_UNKNOWN,
				Value: "bcd",
				Start: TokenPosition{Line: 0, Column: 1},
				End:   TokenPosition{Line: 0, Column: 4},
			},
			currentTokenStart: 1,
			position:          4,
		},
		expectedLexer: &Lexer{
			input:              "abcde",
			tokens:             make(chan Token, 5),
			nextLexingFunction: nil,
			currentToken: Token{
				Type:  TOKEN_UNKNOWN,
				Value: "",
				Start: TokenPosition{Line: 0, Column: 4},
				End:   TokenPosition{Line: 0, Column: 4},
			},
			currentTokenStart: 4,
			position:          4,
		},
	},
	{
		name:           "newline",
		inputTokenType: TOKEN_NEWLINE,
		expectedToken: Token{
			Type:  TOKEN_NEWLINE,
			Value: string(RUNE_NEWLINE),
			Start: TokenPosition{Line: 0, Column: 5},
			End:   TokenPosition{Line: 0, Column: 6},
		},
		inputLexer: &Lexer{
			input:              "abcde\nfghij",
			tokens:             make(chan Token, 5),
			nextLexingFunction: nil,
			currentToken: Token{
				Type:  TOKEN_UNKNOWN,
				Value: "",
				Start: TokenPosition{Line: 0, Column: 5},
				End:   TokenPosition{Line: 0, Column: 6},
			},
			currentTokenStart: 5,
			position:          6,
		},
		expectedLexer: &Lexer{
			input:              "abcde\nfghij",
			tokens:             make(chan Token, 5),
			nextLexingFunction: nil,
			currentToken: Token{
				Type:  TOKEN_UNKNOWN,
				Value: "",
				Start: TokenPosition{Line: 1, Column: 0},
				End:   TokenPosition{Line: 1, Column: 0},
			},
			currentTokenStart: 6,
			position:          6,
		},
	},
	{
		name:           "current token start position out of file",
		inputTokenType: TOKEN_LYRICS,
		expectedToken: Token{
			Type:  TOKEN_ERROR,
			Value: LEXER_ERROR_START_OF_TOKEN_AFTER_EOF,
			Start: TokenPosition{Line: 0, Column: 0},
			End:   TokenPosition{Line: 0, Column: 5},
		},
		inputLexer: &Lexer{
			input:              "abcde",
			tokens:             make(chan Token, 5),
			nextLexingFunction: nil,
			currentToken: Token{
				Type:  TOKEN_UNKNOWN,
				Value: "abcde",
				Start: TokenPosition{Line: 0, Column: 0},
				End:   TokenPosition{Line: 0, Column: 5},
			},
			currentTokenStart: 6,
			position:          5,
		},
		expectedLexer: &Lexer{
			input:              "abcde",
			tokens:             make(chan Token, 5),
			nextLexingFunction: nil,
			currentToken: Token{
				Type:  TOKEN_ERROR,
				Value: LEXER_ERROR_START_OF_TOKEN_AFTER_EOF,
				Start: TokenPosition{Line: 0, Column: 0},
				End:   TokenPosition{Line: 0, Column: 5},
			},
			currentTokenStart: 6,
			position:          5,
		},
	},
	{
		name:           "lexer position out of file",
		inputTokenType: TOKEN_LYRICS,
		expectedToken: Token{
			Type:  TOKEN_ERROR,
			Value: LEXER_ERROR_POSITION_AFTER_EOF,
			Start: TokenPosition{Line: 0, Column: 0},
			End:   TokenPosition{Line: 0, Column: 5},
		},
		inputLexer: &Lexer{
			input:              "abcde",
			tokens:             make(chan Token, 5),
			nextLexingFunction: nil,
			currentToken: Token{
				Type:  TOKEN_UNKNOWN,
				Value: "abcde",
				Start: TokenPosition{Line: 0, Column: 0},
				End:   TokenPosition{Line: 0, Column: 5},
			},
			currentTokenStart: 0,
			position:          6,
		},
		expectedLexer: &Lexer{
			input:              "abcde",
			tokens:             make(chan Token, 5),
			nextLexingFunction: nil,
			currentToken: Token{
				Type:  TOKEN_ERROR,
				Value: LEXER_ERROR_POSITION_AFTER_EOF,
				Start: TokenPosition{Line: 0, Column: 0},
				End:   TokenPosition{Line: 0, Column: 5},
			},
			currentTokenStart: 0,
			position:          6,
		},
	},
}

func TestPushToken(t *testing.T) {
	for _, testCase := range pushTokenTestCases {
		t.Run(testCase.name, func(t *testing.T) {

			testCase.inputLexer.PushToken(testCase.inputTokenType)

			assertEqualLexer(t, testCase.expectedLexer, testCase.inputLexer)

			select {
			case actualToken := <-testCase.inputLexer.tokens:
				assert.Equal(t, testCase.expectedToken, actualToken)
			default:
				t.Fatal("token channel should contain a token")
			}

			select {
			case <-testCase.inputLexer.tokens:
				t.Fatal("no more token should be present on channel")
			default:
			}
		})
	}
}

type nextTokenTestCase struct {
	name              string
	expectedToken     Token
	expectedCallCount int
	inputLexer        *Lexer
	expectedLexer     *Lexer
}

var nextTokenTestCase_tokenInChannel = nextTokenTestCase{
	name:              "token in token channel",
	expectedToken:     testToken,
	expectedCallCount: 0,
	inputLexer: &Lexer{
		input:              "abcde",
		tokens:             make(chan Token, 5),
		nextLexingFunction: mockLexingFunction,
		currentToken: Token{
			Type:  TOKEN_UNKNOWN,
			Value: "bcd",
			Start: TokenPosition{Line: 0, Column: 1},
			End:   TokenPosition{Line: 0, Column: 4},
		},
		currentTokenStart: 1,
		position:          4,
	},
	expectedLexer: &Lexer{
		input:              "abcde",
		tokens:             make(chan Token, 5),
		nextLexingFunction: mockLexingFunction,
		currentToken: Token{
			Type:  TOKEN_UNKNOWN,
			Value: "bcd",
			Start: TokenPosition{Line: 0, Column: 1},
			End:   TokenPosition{Line: 0, Column: 4},
		},
		currentTokenStart: 1,
		position:          4,
	},
}
var nextTokenTestCase_noTokenInChannel = nextTokenTestCase{
	name:              "no token in token channel",
	expectedToken:     testToken,
	expectedCallCount: 1,
	inputLexer: &Lexer{
		input:              "abcde",
		tokens:             make(chan Token, 5),
		nextLexingFunction: mockLexingFunction,
		currentToken: Token{
			Type:  TOKEN_UNKNOWN,
			Value: "bcd",
			Start: TokenPosition{Line: 0, Column: 1},
			End:   TokenPosition{Line: 0, Column: 4},
		},
		currentTokenStart: 1,
		position:          4,
	},
	expectedLexer: &Lexer{
		input:              "abcde",
		tokens:             make(chan Token, 5),
		nextLexingFunction: mockLexingFunction,
		currentToken: Token{
			Type:  TOKEN_UNKNOWN,
			Value: "bcd",
			Start: TokenPosition{Line: 0, Column: 1},
			End:   TokenPosition{Line: 0, Column: 4},
		},
		currentTokenStart: 1,
		position:          4,
	},
}
var nextTokenTestCase_nilLexingFunc = nextTokenTestCase{
	name: "nil lexing function",
	expectedToken: Token{
		Type:  TOKEN_ERROR,
		Value: LEXER_ERROR_NIL_LEXING_FUNCTION,
		Start: TokenPosition{Line: 0, Column: 1},
		End:   TokenPosition{Line: 0, Column: 4},
	},

	expectedCallCount: 0,
	inputLexer: &Lexer{
		input:              "abcde",
		tokens:             make(chan Token, 5),
		nextLexingFunction: nil,
		currentToken: Token{
			Type:  TOKEN_UNKNOWN,
			Value: "bcd",
			Start: TokenPosition{Line: 0, Column: 1},
			End:   TokenPosition{Line: 0, Column: 4},
		},
		currentTokenStart: 1,
		position:          4,
	},
	expectedLexer: &Lexer{
		input:              "abcde",
		tokens:             make(chan Token, 5),
		nextLexingFunction: mockLexingFunction,
		currentToken: Token{
			Type:  TOKEN_ERROR,
			Value: LEXER_ERROR_NIL_LEXING_FUNCTION,
			Start: TokenPosition{Line: 0, Column: 1},
			End:   TokenPosition{Line: 0, Column: 4},
		},

		currentTokenStart: 1,
		position:          4,
	},
}

func TestNextToken(t *testing.T) {
	t.Run(nextTokenTestCase_tokenInChannel.name, func(t *testing.T) {

		mockLexingFuncCallCounter.ResetCount()

		nextTokenTestCase_tokenInChannel.inputLexer.tokens <- testToken

		actualToken := nextTokenTestCase_tokenInChannel.inputLexer.NextToken()

		assert.Equal(t, nextTokenTestCase_tokenInChannel.expectedToken, actualToken)
		assert.Equal(t, nextTokenTestCase_tokenInChannel.expectedCallCount, mockLexingFuncCallCounter.GetCount())
		assertEqualLexer(t, nextTokenTestCase_tokenInChannel.expectedLexer, nextTokenTestCase_tokenInChannel.inputLexer)
	})

	t.Run(nextTokenTestCase_noTokenInChannel.name, func(t *testing.T) {

		mockLexingFuncCallCounter.ResetCount()

		actualToken := nextTokenTestCase_noTokenInChannel.inputLexer.NextToken()

		assert.Equal(t, nextTokenTestCase_noTokenInChannel.expectedToken, actualToken)
		assert.Equal(t, nextTokenTestCase_noTokenInChannel.expectedCallCount, mockLexingFuncCallCounter.GetCount())
		assertEqualLexer(t, nextTokenTestCase_noTokenInChannel.expectedLexer, nextTokenTestCase_noTokenInChannel.inputLexer)
	})

	t.Run(nextTokenTestCase_nilLexingFunc.name, func(t *testing.T) {

		mockLexingFuncCallCounter.ResetCount()

		actualToken := nextTokenTestCase_nilLexingFunc.inputLexer.NextToken()

		assert.Equal(t, nextTokenTestCase_nilLexingFunc.expectedToken, actualToken)
		assert.Equal(t, nextTokenTestCase_nilLexingFunc.expectedCallCount, mockLexingFuncCallCounter.GetCount())
		assertEqualLexer(t, nextTokenTestCase_nilLexingFunc.expectedLexer, nextTokenTestCase_nilLexingFunc.inputLexer)
	})
}

var errorfTestCase = struct {
	inputFormatMessage string
	inputArgument      interface{}
	expectedCallCount  int
	inputLexer         *Lexer
	expectedLexer      *Lexer
}{
	inputFormatMessage: "test message: %d",
	inputArgument:      12,
	expectedCallCount:  0,
	inputLexer: &Lexer{
		input:              "abcde",
		tokens:             make(chan Token, 5),
		nextLexingFunction: nil,
		currentToken: Token{
			Type:  TOKEN_UNKNOWN,
			Value: "bcd",
			Start: TokenPosition{Line: 0, Column: 1},
			End:   TokenPosition{Line: 0, Column: 4},
		},
		currentTokenStart: 1,
		position:          4,
	},
	expectedLexer: &Lexer{
		input:              "abcde",
		tokens:             make(chan Token, 5),
		nextLexingFunction: nil,
		currentToken: Token{
			Type:  TOKEN_ERROR,
			Value: "test message: 12",
			Start: TokenPosition{Line: 0, Column: 1},
			End:   TokenPosition{Line: 0, Column: 4},
		},

		currentTokenStart: 1,
		position:          4,
	},
}

func TestErrorf(t *testing.T) {
	mockLexingFuncCallCounter.ResetCount()

	actualLexingFunction := errorfTestCase.inputLexer.Errorf(errorfTestCase.inputFormatMessage, errorfTestCase.inputArgument)

	assert.Nil(t, actualLexingFunction)
	assert.Equal(t, errorfTestCase.expectedCallCount, mockLexingFuncCallCounter.GetCount())
	assertEqualLexer(t, errorfTestCase.expectedLexer, errorfTestCase.inputLexer)

}

var skipWhitespaceTestCase_skipWhitespaceWithNewline = struct {
	name              string
	expectedCallCount int
	inputLexer        *Lexer
	expectedLexer     *Lexer
}{
	name:              "skip whitespaces with newline",
	expectedCallCount: 0,
	inputLexer: &Lexer{
		input:              " \t\r\nabcde",
		tokens:             make(chan Token, 5),
		nextLexingFunction: nil,
		currentToken: Token{
			Type:  TOKEN_UNKNOWN,
			Value: "",
			Start: TokenPosition{Line: 0, Column: 0},
			End:   TokenPosition{Line: 0, Column: 0},
		},
		currentTokenStart: 0,
		position:          0,
	},
	expectedLexer: &Lexer{
		input:              " \t\r\nabcde",
		tokens:             make(chan Token, 5),
		nextLexingFunction: nil,
		currentToken: Token{
			Type:  TOKEN_UNKNOWN,
			Value: "",
			Start: TokenPosition{Line: 1, Column: 0},
			End:   TokenPosition{Line: 1, Column: 0},
		},
		currentTokenStart: 4,
		position:          4,
	},
}

var skipWhitespaceTestCase_EOF = struct {
	name              string
	expectedCallCount int
	expectedToken     Token
	inputLexer        *Lexer
	expectedLexer     *Lexer
}{
	name: "EOF",
	expectedToken: Token{
		Type:  TOKEN_EOF,
		Value: "abcde \n",
		Start: TokenPosition{Line: 1, Column: 0},
		End:   TokenPosition{Line: 1, Column: 0},
	},
	inputLexer: &Lexer{
		input:              "abcde \n",
		tokens:             make(chan Token, 5),
		nextLexingFunction: nil,
		currentToken: Token{
			Type:  TOKEN_UNKNOWN,
			Value: "abcde",
			Start: TokenPosition{Line: 0, Column: 0},
			End:   TokenPosition{Line: 0, Column: 5},
		},
		currentTokenStart: 0,
		position:          5,
	},
	expectedLexer: &Lexer{
		input:              "abcde \n",
		tokens:             make(chan Token, 5),
		nextLexingFunction: nil,
		currentToken: Token{
			Type:  TOKEN_UNKNOWN,
			Value: "",
			Start: TokenPosition{Line: 1, Column: 0},
			End:   TokenPosition{Line: 1, Column: 0},
		},
		currentTokenStart: 7,
		position:          7,
	},
}

func TestSkipWitheSpace(t *testing.T) {
	t.Run(skipWhitespaceTestCase_skipWhitespaceWithNewline.name, func(t *testing.T) {
		mockLexingFuncCallCounter.ResetCount()

		skipWhitespaceTestCase_skipWhitespaceWithNewline.inputLexer.SkipWhitespace()

		assert.Equal(t, skipWhitespaceTestCase_skipWhitespaceWithNewline.expectedCallCount, mockLexingFuncCallCounter.GetCount())
		assertEqualLexer(t, skipWhitespaceTestCase_skipWhitespaceWithNewline.expectedLexer, skipWhitespaceTestCase_skipWhitespaceWithNewline.inputLexer)
	})
	t.Run(skipWhitespaceTestCase_EOF.name, func(t *testing.T) {
		mockLexingFuncCallCounter.ResetCount()

		skipWhitespaceTestCase_EOF.inputLexer.SkipWhitespace()

		assert.Equal(t, skipWhitespaceTestCase_EOF.expectedCallCount, mockLexingFuncCallCounter.GetCount())
		assertEqualLexer(t, skipWhitespaceTestCase_EOF.expectedLexer, skipWhitespaceTestCase_EOF.inputLexer)

		select {
		case actualToken := <-skipWhitespaceTestCase_EOF.inputLexer.tokens:
			assert.Equal(t, skipWhitespaceTestCase_EOF.expectedToken, actualToken)
		default:
			t.Fatal("token channel should contain a token")
		}

		select {
		case <-skipWhitespaceTestCase_EOF.inputLexer.tokens:
			t.Fatal("no more token should be present on channel")
		default:
		}
	})
}

var newlineTestCase = struct {
	expectedCallCount int
	inputLexer        *Lexer
	expectedLexer     *Lexer
}{
	expectedCallCount: 0,
	inputLexer: &Lexer{
		input:              "abcde",
		tokens:             make(chan Token, 5),
		nextLexingFunction: mockLexingFunction,
		currentToken: Token{
			Type:  TOKEN_UNKNOWN,
			Value: "bcd",
			Start: TokenPosition{Line: 0, Column: 1},
			End:   TokenPosition{Line: 0, Column: 4},
		},
		currentTokenStart: 1,
		position:          4,
	},
	expectedLexer: &Lexer{
		input:              "abcde",
		tokens:             make(chan Token, 5),
		nextLexingFunction: mockLexingFunction,
		currentToken: Token{
			Type:  TOKEN_UNKNOWN,
			Value: "bcd",
			Start: TokenPosition{Line: 1, Column: 0},
			End:   TokenPosition{Line: 1, Column: 0},
		},

		currentTokenStart: 1,
		position:          4,
	},
}

func TestNewLine(t *testing.T) {
	mockLexingFuncCallCounter.ResetCount()

	newlineTestCase.inputLexer.countNewline()

	assert.Equal(t, newlineTestCase.expectedCallCount, mockLexingFuncCallCounter.GetCount())
	assertEqualLexer(t, newlineTestCase.expectedLexer, newlineTestCase.inputLexer)
}

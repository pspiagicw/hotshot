package lexer

import (
	"testing"

	"github.com/pspiagicw/hotshot/token"
)

func TestQuote(t *testing.T) {
	input := `
    'something
    (type 'something)
    `

	expectedTokens := []token.Token{
		{TokenType: token.QUOTE, TokenValue: "'"},
		{TokenType: token.IDENT, TokenValue: "something"},
		{TokenType: token.LPAREN, TokenValue: "("},
		{TokenType: token.IDENT, TokenValue: "type"},
		{TokenType: token.QUOTE, TokenValue: "'"},
		{TokenType: token.IDENT, TokenValue: "something"},
		{TokenType: token.RPAREN, TokenValue: ")"},
		{TokenType: token.EOF, TokenValue: " "},
	}

	checkTokens(t, expectedTokens, input)
}
func TestIdent(t *testing.T) {

	input := "(+ one two)"

	expectedTokens := []token.Token{
		{TokenType: token.LPAREN, TokenValue: "("},
		{TokenType: token.PLUS, TokenValue: "+"},
		{TokenType: token.IDENT, TokenValue: "one"},
		{TokenType: token.IDENT, TokenValue: "two"},
		{TokenType: token.RPAREN, TokenValue: ")"},
		{TokenType: token.EOF, TokenValue: " "},
	}
	checkTokens(t, expectedTokens, input)

}
func TestIdent2(t *testing.T) {
	input := "(one)"

	expectedTokens := []token.Token{
		{TokenType: token.LPAREN, TokenValue: "("},
		{TokenType: token.IDENT, TokenValue: "one"},
		{TokenType: token.RPAREN, TokenValue: ")"},
		{TokenType: token.EOF, TokenValue: " "},
	}
	checkTokens(t, expectedTokens, input)
}

func TestParen(t *testing.T) {
	input := "(){}[]"

	expectedTokens := []token.Token{
		{TokenType: token.LPAREN, TokenValue: "("},
		{TokenType: token.RPAREN, TokenValue: ")"},
		{TokenType: token.LBRACE, TokenValue: "{"},
		{TokenType: token.RBRACE, TokenValue: "}"},
		{TokenType: token.LSQUARE, TokenValue: "["},
		{TokenType: token.RSQUARE, TokenValue: "]"},
		{TokenType: token.EOF, TokenValue: " "},
	}
	checkTokens(t, expectedTokens, input)
}

func TestParenWithSpaces(t *testing.T) {
	input := "(     )"

	expectedTokens := []token.Token{
		{TokenType: token.LPAREN, TokenValue: "("},
		{TokenType: token.RPAREN, TokenValue: ")"},
		{TokenType: token.EOF, TokenValue: " "},
	}
	checkTokens(t, expectedTokens, input)
}
func TestStrings(t *testing.T) {
	input := `"something"`

	expectedTokens := []token.Token{
		{TokenType: token.STRING, TokenValue: "something"},
		{TokenType: token.EOF, TokenValue: " "},
	}
	checkTokens(t, expectedTokens, input)
}
func TestComment(t *testing.T) {
	input := "; this is a comment ;"

	expectedTokens := []token.Token{
		{TokenType: token.EOF, TokenValue: " "},
	}
	checkTokens(t, expectedTokens, input)
}
func TestKeywords(t *testing.T) {

	input := "if true false while case fn lambda cond let assert import set"

	expectedTokens := []token.Token{
		{TokenType: token.IF, TokenValue: "if"},
		{TokenType: token.TRUE, TokenValue: "true"},
		{TokenType: token.FALSE, TokenValue: "false"},
		{TokenType: token.WHILE, TokenValue: "while"},
		{TokenType: token.CASE, TokenValue: "case"},
		{TokenType: token.FN, TokenValue: "fn"},
		{TokenType: token.LAMBDA, TokenValue: "lambda"},
		{TokenType: token.COND, TokenValue: "cond"},
		{TokenType: token.LET, TokenValue: "let"},
		{TokenType: token.ASSERT, TokenValue: "assert"},
		{TokenType: token.IMPORT, TokenValue: "import"},
		{TokenType: token.SET, TokenValue: "set"},
		{TokenType: token.EOF, TokenValue: " "},
	}
	checkTokens(t, expectedTokens, input)

}
func TestNum2(t *testing.T) {
	input := "(1)"

	expectedTokens := []token.Token{
		{TokenType: token.LPAREN, TokenValue: "("},
		{TokenType: token.NUM, TokenValue: "1"},
		{TokenType: token.RPAREN, TokenValue: ")"},
		{TokenType: token.EOF, TokenValue: " "},
	}

	checkTokens(t, expectedTokens, input)

}

func TestNumNegative(t *testing.T) {

	input := "(-1)+1"

	expectedTokens := []token.Token{
		{TokenType: token.LPAREN, TokenValue: "("},
		{TokenType: token.NUM, TokenValue: "-1"},
		{TokenType: token.RPAREN, TokenValue: ")"},
		{TokenType: token.NUM, TokenValue: "1"},
		{TokenType: token.EOF, TokenValue: " "},
	}

	checkTokens(t, expectedTokens, input)
}
func TestNumbers(t *testing.T) {

	input := "1"

	expectedTokens := []token.Token{
		{TokenType: token.NUM, TokenValue: "1"},
		{TokenType: token.EOF, TokenValue: " "},
	}

	checkTokens(t, expectedTokens, input)

}

func TestSymbols(t *testing.T) {

	input := "+-/*^"

	expectedTokens := []token.Token{
		{TokenType: token.PLUS, TokenValue: "+"},
		{TokenType: token.MINUS, TokenValue: "-"},
		{TokenType: token.SLASH, TokenValue: "/"},
		{TokenType: token.MULTIPLY, TokenValue: "*"},
		{TokenType: token.POWER, TokenValue: "^"},
		{TokenType: token.EOF, TokenValue: " "},
	}
	checkTokens(t, expectedTokens, input)

}

func TestComparator(t *testing.T) {

	input := "<>="

	expectedTokens := []token.Token{
		{TokenType: token.LESSTHAN, TokenValue: "<"},
		{TokenType: token.GREATERTHAN, TokenValue: ">"},
		{TokenType: token.EQ, TokenValue: "="},
		{TokenType: token.EOF, TokenValue: " "},
	}
	checkTokens(t, expectedTokens, input)

}
func TestOthers(t *testing.T) {
	input := "|^"

	expectedTokens := []token.Token{
		{TokenType: token.PIPE, TokenValue: "|"},
		{TokenType: token.POWER, TokenValue: "^"},
		{TokenType: token.EOF, TokenValue: " "},
	}
	checkTokens(t, expectedTokens, input)

}
func checkTokens(t *testing.T, expected []token.Token, input string) {
	t.Helper()

	lexer := NewLexer(input)

	for i, expectedToken := range expected {
		actualToken := lexer.Next()
		matchToken(t, i, expectedToken, actualToken)
	}

}
func matchToken(t *testing.T, i int, expected token.Token, actual *token.Token) {
	t.Helper()
	if actual.TokenType != expected.TokenType {
		t.Errorf("Test [%d], Expected TokenType: '%v', Actual TokenType: '%v'", i, expected.TokenType, actual.TokenType)

	}
	if actual.TokenValue != expected.TokenValue {
		t.Errorf("Test [%d], Expected TokenValue: '%v', Actual TokenValue: '%v'", i, expected.TokenValue, actual.TokenValue)
	}

}

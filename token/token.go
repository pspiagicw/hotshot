package token

import "fmt"

type TokenType string

type Token struct {
	TokenType  TokenType
	TokenValue string
}

func (t *Token) String() string {
	return fmt.Sprintf("Token(Type: %s, Value: '%s')", t.TokenType, t.TokenValue)
}

const (
	LPAREN = "LPAREN"
	RPAREN = "RPAREN"

	PLUS     = "PLUS"
	MINUS    = "MINUS"
	MULTIPLY = "MULTIPLY"
	SLASH    = "SLASH"
	MOD      = "MOD"

	LESSTHAN    = "LESSTHAN"
	GREATERTHAN = "GREATERTHAN"
	EQ          = "EQ"

	IDENT = "IDENT"
	NUM   = "NUM"

	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	STRING    = "STRING"
	PIPE      = "PIPE"
	SEMICOLON = "SEMICOLON"

	DOLLAR   = "DOLLAR"
	AT       = "AT"
	BANG     = "BANG"
	COMMA    = "COMMA"
	QUESTION = "QUESTION"
)

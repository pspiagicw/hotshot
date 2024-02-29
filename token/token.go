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
	COMMENT = "COMMENT"

	STRING = "STRING"
	PIPE   = "PIPE"

	DOLLAR   = "DOLLAR"
	AT       = "AT"
	BANG     = "BANG"
	COMMA    = "COMMA"
	QUESTION = "QUESTION"
	HASH     = "HASH"
	POWER    = "POWER"

	IF    = "IF"
	TRUE  = "TRUE"
	FALSE = "FALSE"
	WHILE = "WHILE"
	CASE  = "CASE"
	FN    = "FN"
)

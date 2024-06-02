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
	LPAREN  = "LPAREN"
	RPAREN  = "RPAREN"
	LBRACE  = "LBRACE"
	RBRACE  = "RBRACE"
	LSQUARE = "LSQUARE"
	RSQUARE = "RSQUARE"

	PLUS     = "PLUS"
	MINUS    = "MINUS"
	MULTIPLY = "MULTIPLY"
	SLASH    = "SLASH"

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

	POWER = "POWER"

	IF     = "IF"
	TRUE   = "TRUE"
	FALSE  = "FALSE"
	WHILE  = "WHILE"
	CASE   = "CASE"
	FN     = "FN"
	LAMBDA = "LAMBDA"
	COND   = "COND"
	LET    = "LET"
	ASSERT = "ASSERT"
	IMPORT = "IMPORT"
	SET    = "SET"

	QUOTE = "QUOTE"
)

package token

type TokenType string

type Token struct {
	TokenType  TokenType
	TokenValue byte
}

// var LPAREN string = "LPAREN"
const (
	LPAREN  = "LPAREN"
	RPAREN  = "RPAREN"
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"
)

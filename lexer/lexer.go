package lexer

import "github.com/pspiagicw/hotshot/token"

type Lexer struct {
	input       string
	curPos      int
	readPos     int
	currentChar byte
	EOF         bool
	inputLen    int
}

func createEOF() *token.Token {
	return &token.Token{
		TokenType:  token.EOF,
		TokenValue: ' ',
	}

}
func createILLEGAL() *token.Token {
	return &token.Token{
		TokenType:  token.ILLEGAL,
		TokenValue: ' ',
	}
}

func (l *Lexer) consumeSpaces() {
	for l.currentChar == ' ' || l.currentChar == '\t' {
		l.advance()
	}
}
func (l *Lexer) Next() *token.Token {
	returnToken := createILLEGAL()

	if l.EOF {
		return createEOF()
	}

	l.consumeSpaces()

	switch l.currentChar {
	case '(':
		returnToken.TokenType = token.LPAREN
		returnToken.TokenValue = l.currentChar
	case ')':
		returnToken.TokenType = token.RPAREN
		returnToken.TokenValue = l.currentChar
	}
	l.advance()

	return returnToken

}
func (l *Lexer) advance() {
	l.curPos = l.readPos
	l.readPos += 1

	if l.readPos > l.inputLen {
		l.EOF = true
	} else {
		l.currentChar = l.input[l.curPos]
	}
}

func NewLexer(input string) *Lexer {
	l := &Lexer{
		input:       input,
		curPos:      -1,
		readPos:     0,
		inputLen:    len(input),
		EOF:         false,
		currentChar: ' ',
	}
	return l
}

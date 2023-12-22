package lexer

import (
	"unicode"

	"github.com/pspiagicw/hotshot/token"
)

type Lexer struct {
	input       string
	curPos      int
	readPos     int
	currentChar string
	EOF         bool
	inputLen    int
}

func createEOF() *token.Token {
	return &token.Token{
		TokenType:  token.EOF,
		TokenValue: " ",
	}

}
func (l *Lexer) createILLEGAL() *token.Token {
	return &token.Token{
		TokenType:  token.ILLEGAL,
		TokenValue: l.currentChar,
	}
}

func (l *Lexer) consumeSpaces() {
	for l.currentChar == " " || l.currentChar == "\t" || l.currentChar == "\n" {
		l.advance()
	}
}
func (l *Lexer) peekInput() string {
	if l.readPos < l.inputLen {
		return string(l.input[l.readPos])

	}
	return " "

}
func (l *Lexer) Next() *token.Token {

	shouldAdvance := true

	l.consumeSpaces()
	if l.readPos < 0 {
		l.advance()
	}

	returnToken := l.createILLEGAL()

	if l.EOF {
		return createEOF()
	}

	switch l.currentChar {
	case "(":
		returnToken.TokenType = token.LPAREN
		returnToken.TokenValue = l.currentChar
	case ")":
		returnToken.TokenType = token.RPAREN
		returnToken.TokenValue = l.currentChar
	case "|":
		returnToken.TokenType = token.PIPE
		returnToken.TokenValue = l.currentChar
	case "@":
		returnToken.TokenType = token.AT
		returnToken.TokenValue = l.currentChar
	case ";":
		returnToken.TokenType = token.SEMICOLON
		returnToken.TokenValue = l.currentChar
	case "$":
		returnToken.TokenType = token.DOLLAR
		returnToken.TokenValue = l.currentChar
	case "!":
		returnToken.TokenType = token.BANG
		returnToken.TokenValue = l.currentChar
	case ",":
		returnToken.TokenType = token.COMMA
		returnToken.TokenValue = l.currentChar
	case "?":
		returnToken.TokenType = token.QUESTION
		returnToken.TokenValue = l.currentChar
	case "+":
		if l.isDigit(l.peekInput()) {
			l.advance()

			number := l.extractNumber()

			returnToken.TokenType = token.NUM
			returnToken.TokenValue = number

			shouldAdvance = false
		} else {
			returnToken.TokenType = token.PLUS
			returnToken.TokenValue = l.currentChar
		}
	case "-":
		if l.isDigit(l.peekInput()) {
			l.advance()

			number := l.extractNumber()

			returnToken.TokenType = token.NUM
			returnToken.TokenValue = "-" + number

			shouldAdvance = false

		} else {
			returnToken.TokenType = token.MINUS
			returnToken.TokenValue = l.currentChar
		}
	case "/":
		returnToken.TokenType = token.SLASH
		returnToken.TokenValue = l.currentChar
	case "*":
		returnToken.TokenType = token.MULTIPLY
		returnToken.TokenValue = l.currentChar
	case "=":
		returnToken.TokenType = token.EQ
		returnToken.TokenValue = l.currentChar
	case "%":
		returnToken.TokenType = token.MOD
		returnToken.TokenValue = l.currentChar
	case "<":
		returnToken.TokenType = token.LESSTHAN
		returnToken.TokenValue = l.currentChar
	case ">":
		returnToken.TokenType = token.GREATERTHAN
		returnToken.TokenValue = l.currentChar
	case "\"":
		l.advance()
		string := l.extractString()
		returnToken.TokenType = token.STRING
		returnToken.TokenValue = string
	default:
		if l.isLetter(l.currentChar) {
			identifier := l.extractIdentifier()

			returnToken.TokenType = token.IDENT
			returnToken.TokenValue = identifier

			shouldAdvance = false

		} else if l.isDigit(l.currentChar) {
			number := l.extractNumber()

			returnToken.TokenType = token.NUM
			returnToken.TokenValue = number

			shouldAdvance = false

		}

	}

	if shouldAdvance {
		l.advance()
	}

	return returnToken

}
func (l *Lexer) extractIdentifier() string {
	identifier := ""
	for l.isLetter(l.currentChar) {
		identifier += l.currentChar
		l.advance()
	}
	return identifier
}
func (l *Lexer) extractString() string {
	identifier := ""
	for l.currentChar != "\"" {
		identifier += l.currentChar
		l.advance()
	}
	return identifier
}

func (l *Lexer) extractNumber() string {
	number := ""
	for l.isDigit(l.currentChar) && !l.EOF {
		// fmt.Println(l.input, l.currentChar, l.curPos, l.readPos)
		number += l.currentChar
		l.advance()
	}
	return number
}

func (l *Lexer) isLetter(input string) bool {
	if l.EOF {
		return false
	}

	char := input[0]

	return !l.EOF && unicode.IsLetter(rune(char))

}
func (l *Lexer) isDigit(input string) bool {
	if l.EOF {
		return false
	}

	if '0' <= input[0] && input[0] <= '9' {
		return true
	}
	return false
}
func (l *Lexer) advance() {
	l.curPos = l.readPos
	l.readPos += 1

	if l.readPos > l.inputLen {
		l.EOF = true
		l.currentChar = ""
	} else {
		l.currentChar = string(l.input[l.curPos])
	}
}

func NewLexer(input string) *Lexer {
	l := &Lexer{
		input:       input,
		curPos:      -1,
		readPos:     0,
		inputLen:    len(input),
		EOF:         false,
		currentChar: " ",
	}
	return l
}

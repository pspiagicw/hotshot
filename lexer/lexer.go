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
	for l.currentChar == " " || l.currentChar == "\t" || l.currentChar == "\n" && l.EOF == false {
		l.advance()
	}
}
func (l *Lexer) consumeComments() {
	if l.currentChar == ";" {
		l.advance()
		for l.currentChar != ";" && l.EOF == false {
			l.advance()
		}
	}
	if l.currentChar == ";" {
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

	// Consume any start empty lines.
	l.consumeSpaces()
	// Consume any comments.
	l.consumeComments()
	// Consume any newlines after the comments
	l.consumeSpaces()

	if l.readPos < 0 {
		l.advance()
	}

	returnToken := l.createILLEGAL()

	if l.EOF {
		return createEOF()
	}

	switch l.currentChar {
	case "'":
		returnToken.TokenType = token.QUOTE
		returnToken.TokenValue = l.currentChar
	case "(":
		returnToken.TokenType = token.LPAREN
		returnToken.TokenValue = l.currentChar
	case ")":
		returnToken.TokenType = token.RPAREN
		returnToken.TokenValue = l.currentChar
	case "{":
		returnToken.TokenType = token.LBRACE
		returnToken.TokenValue = l.currentChar
	case "}":
		returnToken.TokenType = token.RBRACE
		returnToken.TokenValue = l.currentChar
	case "[":
		returnToken.TokenType = token.LSQUARE
		returnToken.TokenValue = l.currentChar
	case "]":
		returnToken.TokenType = token.RSQUARE
		returnToken.TokenValue = l.currentChar
	case "|":
		returnToken.TokenType = token.PIPE
		returnToken.TokenValue = l.currentChar
	case "^":
		returnToken.TokenType = token.POWER
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
	case ";":
		return l.Next()
	default:
		if l.isLetter(l.currentChar) {
			identifier := l.extractIdentifier()

			returnToken = l.parseKeyword(identifier)

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

// func (l *Lexer) extractComment() string {
// 	comment := ""
// 	for l.currentChar != ";" {
// 		if l.EOF {
// 			return comment
// 		}
// 		comment += l.currentChar
// 		l.advance()
// 	}
// 	return comment
// }

func (l *Lexer) parseKeyword(identifier string) *token.Token {
	keyword := token.Token{
		TokenValue: identifier,
		TokenType:  token.IDENT,
	}
	switch identifier {
	case "case":
		keyword.TokenType = token.CASE
	case "if":
		keyword.TokenType = token.IF
	case "true":
		keyword.TokenType = token.TRUE
	case "false":
		keyword.TokenType = token.FALSE
	case "while":
		keyword.TokenType = token.WHILE
	case "fn":
		keyword.TokenType = token.FN
	case "lambda":
		keyword.TokenType = token.LAMBDA
	case "cond":
		keyword.TokenType = token.COND
	case "let":
		keyword.TokenType = token.LET
	case "assert":
		keyword.TokenType = token.ASSERT
	case "import":
		keyword.TokenType = token.IMPORT
	case "set":
		keyword.TokenType = token.SET
	}
	return &keyword
}
func (l *Lexer) extractIdentifier() string {
	identifier := ""
	for l.isLetter(l.currentChar) || l.currentChar == "-" || l.isDigit(l.currentChar) {
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

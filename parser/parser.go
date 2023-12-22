package parser

import (
	"fmt"

	"github.com/pspiagicw/hotshot/ast"
	"github.com/pspiagicw/hotshot/lexer"
	"github.com/pspiagicw/hotshot/token"
)

type Parser struct {
	l         *lexer.Lexer
	errors    []error
	curToken  *token.Token
	peekToken *token.Token
}

func NewParser(l *lexer.Lexer) *Parser {
	return &Parser{
		l:         l,
		errors:    []error{},
		peekToken: l.Next(),
	}
}
func (p *Parser) advance() {
	p.curToken = p.peekToken
	p.peekToken = p.l.Next()
}

func (p *Parser) Parse() *ast.Program {
	program := new(ast.Program)
	for p.curToken == nil || p.curToken.TokenType != token.EOF {
		currentStatement := p.parseStatement()
		program.Statements = append(program.Statements, currentStatement)
	}
	return program
}

func (p *Parser) parseStatement() ast.Statement {
	p.advance()
	switch p.curToken.TokenType {
	case token.NUM:
		return p.parseIntStatement()
	case token.STRING:
		return p.parseStringStatement()
	case token.LPAREN:
		return p.parseComplexStatement()
	case token.ILLEGAL:
		p.registerError(fmt.Errorf("Expected token found: %v", p.curToken.String()))
	}
	return nil
}

func (p *Parser) parseComplexStatement() ast.Statement {
	p.advance()
	switch p.curToken.TokenType {
	case token.RPAREN:
		return &ast.EmptyStatement{}
	default:
		return p.parseFunctionCall()
	}
}

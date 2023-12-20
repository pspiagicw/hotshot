package parser

import (
	"fmt"
	"strconv"

	"github.com/pspiagicw/hotshot/ast"
	"github.com/pspiagicw/hotshot/lexer"
	"github.com/pspiagicw/hotshot/token"
)

type Parser struct {
	l      *lexer.Lexer
	errors []error
}

func NewParser(l *lexer.Lexer) *Parser {
	return &Parser{
		l:      l,
		errors: []error{},
	}
}

func (p *Parser) Parse() *ast.Program {
	program := new(ast.Program)
	for !p.l.EOF {
		currentStatement := p.parseStatement()
		program.Statements = append(program.Statements, currentStatement)
	}
	return program
}

func (p *Parser) registerError(err error) {
	p.errors = append(p.errors, err)
}

func (p *Parser) Errors() []error {
	return p.errors
}

func (p *Parser) parseStatement() ast.Statement {
	identifierToken := p.l.Next()
	switch identifierToken.TokenType {
	case token.NUM:
		return p.parseIntStatement(identifierToken)
	case token.LPAREN:
		return p.parseConcreteStatement(identifierToken)
	}
	return nil
}
func (p *Parser) parseIntStatement(start *token.Token) *ast.IntStatement {
	value, err := strconv.Atoi(start.TokenValue)
	if err != nil {
		p.registerError(fmt.Errorf("Error parsing string into Integer: %v", err))
	}
	return &ast.IntStatement{
		Value: value,
	}
}
func (p *Parser) parseConcreteStatement(identifierToken *token.Token) ast.Statement {
	unitToken := p.l.Next()
	for unitToken.TokenType != token.RPAREN {
		unitToken = p.l.Next()
	}
	return &ast.EmptyStatement{}

}

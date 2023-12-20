package parser

import (
	"fmt"
	"strconv"

	"github.com/pspiagicw/hotshot/ast"
	"github.com/pspiagicw/hotshot/lexer"
	"github.com/pspiagicw/hotshot/token"
)

type Parser struct {
	l        *lexer.Lexer
	errors   []error
	curToken *token.Token
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
		if p.curToken.TokenType != token.EOF {
			program.Statements = append(program.Statements, currentStatement)
		}
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
	start := p.l.Next()
	p.curToken = start
	switch start.TokenType {
	case token.NUM:
		return p.parseIntStatement(start)
	case token.STRING:
		return p.parseStringStatement(start)
	case token.LPAREN:
		return p.parseConcreteStatement(start)
	}
	return nil
}
func (p *Parser) parseStringStatement(start *token.Token) *ast.StringStatement {
	return &ast.StringStatement{
		Value: start.TokenValue,
	}
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
func (p *Parser) parseConcreteStatement(start *token.Token) ast.Statement {
	unitToken := p.l.Next()
	if unitToken.TokenType == token.RPAREN {
		return &ast.EmptyStatement{}
	} else {
		st := new(ast.FunctionalStatement)
		st.Op = unitToken
		for p.curToken.TokenType != token.RPAREN {
			st.Args = append(st.Args, p.parseStatement())
		}
		return st
	}
}

package parser

import (
	"fmt"
	"os"
	"strconv"

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
	for p.curToken.TokenType != token.EOF {
		if len(p.Errors()) != 0 {
			log.LogError("Error while parsing ")
			for _, err := range p.Errors() {
				log.LogError(err.Error())
			}
			os.Exit(1)
		}
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
	p.advance()
	switch p.curToken.TokenType {
	case token.NUM:
		return p.parseIntStatement()
	case token.STRING:
		return p.parseStringStatement()
	case token.LPAREN:
		return p.parseConcreteStatement()
	}
	return nil
}
func (p *Parser) parseStringStatement() *ast.StringStatement {
	return &ast.StringStatement{
		Value: p.curToken.TokenValue,
	}
}
func (p *Parser) parseIntStatement() *ast.IntStatement {
	value, err := strconv.Atoi(p.curToken.TokenValue)
	if err != nil {
		p.registerError(fmt.Errorf("Error parsing string into Integer: %v", err))
	}
	return &ast.IntStatement{
		Value: value,
	}
}
func (p *Parser) parseConcreteStatement() ast.Statement {
	p.advance()
	if p.curToken.TokenType == token.RPAREN {
		return &ast.EmptyStatement{}
	} else {
		st := new(ast.FunctionalStatement)
		st.Op = p.curToken
		for p.peekToken.TokenType != token.RPAREN {
			if p.peekToken.TokenType == token.EOF || p.peekToken.TokenType == token.ILLEGAL {
				p.registerError(fmt.Errorf("Expected statement, got %s", p.peekToken.TokenType))
			} else {
				st.Args = append(st.Args, p.parseStatement())
			}
		}
		return st
	}
}

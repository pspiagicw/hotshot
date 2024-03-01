package parser

import (
	"fmt"

	"github.com/pspiagicw/hotshot/ast"
	"github.com/pspiagicw/hotshot/lexer"
	"github.com/pspiagicw/hotshot/token"
)

type Parser struct {
	l            *lexer.Lexer
	errors       []error
	curToken     *token.Token
	peekToken    *token.Token
	nilStatement ast.Statement
}

func NewParser(l *lexer.Lexer) *Parser {
	return &Parser{
		l:            l,
		errors:       []error{},
		peekToken:    l.Next(),
		nilStatement: &ast.EmptyStatement{},
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
	case token.TRUE:
		return p.parseBoolStatement()
	case token.FALSE:
		return p.parseBoolStatement()
	case token.ILLEGAL:
		p.registerError(fmt.Errorf("Expected a token for a statement, found: %v", p.curToken.String()))
	case token.IDENT:
		return p.parseIdentStatement()
	case token.EOF:
		return p.nilStatement
	default:
		p.registerError(fmt.Errorf("Expected a token for a statement, found: %v", p.curToken.String()))

	}
	return p.nilStatement
}
func (p *Parser) parseIdentStatement() ast.Statement {
	return &ast.IdentStatement{
		Value: p.curToken,
	}
}
func (p *Parser) parseAssignment() ast.Statement {
	st := &ast.AssignmentStatement{}

	p.expectedTokenIs(token.IDENT)

	st.Name = p.curToken
	st.Value = p.parseStatement()

	p.expectedTokenIs(token.RPAREN)

	return st
}
func (p *Parser) expectedTokenIs(ex token.TokenType) {
	if p.peekToken.TokenType != ex {
		p.registerError(fmt.Errorf("Expected a ident, got %s", p.curToken.TokenType))
	}
	p.advance()
}
func (p *Parser) parseFunctionDec() ast.Statement {
	st := &ast.FunctionStatement{}

	p.expectedTokenIs(token.IDENT)

	st.Name = p.curToken
	p.expectedTokenIs(token.LPAREN)
	st.Args = []ast.Statement{}

	for !p.peekTokenIs(token.RPAREN) {
		arg := p.parseStatement()
		st.Args = append(st.Args, arg)
	}
	p.expectedTokenIs(token.RPAREN)

	st.Body = p.parseStatement()

	p.expectedTokenIs(token.RPAREN)

	return st
}
func (p *Parser) peekTokenIs(ex token.TokenType) bool {
	return p.peekToken.TokenType == ex
}
func (p *Parser) parseIfStatement() ast.Statement {
	st := &ast.IfStatement{}

	st.Condition = p.parseStatement()

	st.Body = p.parseStatement()

	if !p.peekTokenIs(token.RPAREN) {
		st.Else = p.parseStatement()
	}

	// p.advance()
	p.expectedTokenIs(token.RPAREN)

	return st

}
func (p *Parser) parseWhileStatement() ast.Statement {
	st := &ast.WhileStatement{}

	st.Condition = p.parseStatement()

	st.Body = p.parseStatement()

	p.expectedTokenIs(token.RPAREN)

	return st

}

func (p *Parser) parseComplexStatement() ast.Statement {
	p.advance()
	switch p.curToken.TokenType {
	case token.RPAREN:
		return &ast.EmptyStatement{}
	case token.DOLLAR:
		return p.parseAssignment()
	case token.FN:
		return p.parseFunctionDec()
	case token.IF:
		return p.parseIfStatement()
	case token.WHILE:
		return p.parseWhileStatement()
	default:
		return p.parseFunctionCall()
	}
}

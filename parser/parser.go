package parser

import (
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
	case token.LBRACE:
		return p.parseTableStatement()
	case token.ILLEGAL:
		p.registerError("Expected a token for a statement, found: %v", p.curToken.String())
	case token.IDENT:
		return p.parseIdentStatement()
	case token.EOF:
		return p.nilStatement
	default:
		p.registerError("Expected a token for a statement, found: %v", p.curToken.String())
	}
	return p.nilStatement
}
func (p *Parser) parseTableStatement() ast.Statement {
	st := &ast.TableStatement{}

	st.Elements = []ast.Statement{}

	for !p.peekTokenIs(token.RBRACE) {
		element := p.parseStatement()
		st.Elements = append(st.Elements, element)
	}

	p.expectedTokenIs(token.RBRACE)

	return st
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
		p.registerError("Expected a %s, got %s", ex, p.curToken.TokenType)
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
		_, ok := arg.(*ast.IdentStatement)

		if !ok {
			p.registerError("Expected a ident, got %v", arg)
			return nil
		}

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
func (p *Parser) parseLambdaStatement() ast.Statement {
	st := &ast.LambdaStatement{}

	p.expectedTokenIs(token.LPAREN)

	st.Args = []ast.Statement{}

	for !p.peekTokenIs(token.RPAREN) {
		arg := p.parseStatement()
		_, ok := arg.(*ast.IdentStatement)

		if !ok {
			p.registerError("Expected a ident, got %v", arg)
			return nil
		}

		st.Args = append(st.Args, arg)
	}

	p.expectedTokenIs(token.RPAREN)

	st.Body = p.parseStatement()

	p.expectedTokenIs(token.RPAREN)

	return st
}
func (p *Parser) parseCondStatement() ast.Statement {
	st := &ast.CondStatement{}

	// st.Conditions = map[ast.Statement]ast.Statement{}
	st.Expressions = []ast.ConditionExpression{}

	for p.peekTokenIs(token.LPAREN) {
		p.expectedTokenIs(token.LPAREN)
		condition := p.parseStatement()
		body := p.parseStatement()
		p.expectedTokenIs(token.RPAREN)
		st.Expressions = append(st.Expressions, ast.ConditionExpression{
			Condition: condition,
			Body:      body,
		})
	}

	p.expectedTokenIs(token.RPAREN)

	return st
}
func (p *Parser) parseAssertStatement() ast.Statement {
	st := &ast.AssertStatement{}

	st.Left = p.parseStatement()

	st.Right = p.parseStatement()

	p.expectedTokenIs(token.STRING)

	st.Message = p.curToken

	p.expectedTokenIs(token.RPAREN)

	return st

}

func (p *Parser) parseComplexStatement() ast.Statement {
	p.advance()
	switch p.curToken.TokenType {
	case token.RPAREN:
		return &ast.EmptyStatement{}
	case token.LET:
		return p.parseAssignment()
	case token.FN:
		return p.parseFunctionDec()
	case token.IF:
		return p.parseIfStatement()
	case token.WHILE:
		return p.parseWhileStatement()
	case token.LAMBDA:
		return p.parseLambdaStatement()
	case token.COND:
		return p.parseCondStatement()
	case token.ASSERT:
		return p.parseAssertStatement()
	default:
		return p.parseFunctionCall()
	}
}

package parser

import (
	"fmt"
	"strconv"

	"github.com/pspiagicw/hotshot/ast"
	"github.com/pspiagicw/hotshot/token"
)

var validOps = map[token.TokenType]bool{
	token.PLUS:        true,
	token.MINUS:       true,
	token.SLASH:       true,
	token.QUESTION:    true,
	token.HASH:        true,
	token.MULTIPLY:    true,
	token.IF:          true,
	token.CASE:        true,
	token.BANG:        true,
	token.EQ:          true,
	token.GREATERTHAN: true,
	token.LESSTHAN:    true,
	token.MOD:         true,
	token.IDENT:       true,
	token.POWER:       true,
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
func (p *Parser) checkOp(op *token.Token) bool {
	_, ok := validOps[op.TokenType]
	return ok
}

func (p *Parser) parseFunctionCall() *ast.FunctionalStatement {
	st := new(ast.FunctionalStatement)
	if p.checkOp(p.curToken) {
		st.Op = p.curToken
	} else {
		p.registerError(fmt.Errorf("Expected valid op, got %v", p.curToken))
	}

	for p.peekToken.TokenType != token.RPAREN {

		if p.peekToken.TokenType == token.EOF {
			p.registerError(fmt.Errorf("Expected a argument for function call, got %s", p.peekToken.TokenType))
			return nil
		}

		st.Args = append(st.Args, p.parseStatement())
	}

	p.advance()
	return st
}
func (p *Parser) parseBoolStatement() *ast.BoolStatement {
	statement := &ast.BoolStatement{
		Value: p.curToken.TokenType == token.TRUE,
	}
	return statement
}

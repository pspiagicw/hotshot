package parser

import (
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
		p.registerError("Error casting %s as Integer", err)
		return nil
	}
	return &ast.IntStatement{
		Value: value,
	}
}

func (p *Parser) checkOp(op *token.Token) bool {
	_, ok := validOps[op.TokenType]
	return ok
}

func (p *Parser) parseFunctionCall() *ast.CallStatement {
	st := new(ast.CallStatement)
	if p.checkOp(p.curToken) {
		st.Op = p.curToken
	} else {
		p.registerError("Expected a valid function name, got %v", p.curToken)
		return nil
	}

	for p.peekToken.TokenType != token.RPAREN {

		if p.peekToken.TokenType == token.EOF {
			p.registerError("Expected a expression as argument to function OR ), got %s", p.peekToken.TokenType)
			return nil
		}

		st.Args = append(st.Args, p.parseStatement())
	}

	p.expectedTokenIs(token.RPAREN)
	return st
}
func (p *Parser) parseBoolStatement() *ast.BoolStatement {
	statement := &ast.BoolStatement{
		Value: p.curToken.TokenType == token.TRUE,
	}
	return statement
}

package parser

import (
	"fmt"
	"strconv"

	"github.com/pspiagicw/hotshot/ast"
	"github.com/pspiagicw/hotshot/token"
)

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

func (p *Parser) parseFunctionCall() *ast.FunctionalStatement {
	st := new(ast.FunctionalStatement)
	st.Op = p.curToken

	for p.peekToken.TokenType != token.RPAREN {

		if p.peekToken.TokenType == token.EOF {
			p.registerError(fmt.Errorf("Expected token, got %s", p.peekToken.TokenType))
			return nil
		}

		st.Args = append(st.Args, p.parseStatement())
	}

	p.advance()
	return st

}

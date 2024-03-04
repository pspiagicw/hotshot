package parser

import (
	"fmt"
)

func (p *Parser) registerError(format string, v ...interface{}) {
	err := fmt.Errorf(format, v...)
	p.errors = append(p.errors, err)
}

func (p *Parser) Errors() []error {
	return p.errors
}

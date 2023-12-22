package parser

func (p *Parser) registerError(err error) {
	p.errors = append(p.errors, err)
}

func (p *Parser) Errors() []error {
	return p.errors
}

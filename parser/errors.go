package parser

import (
	"fmt"
	"monkey/token"
)

// Errors can be used to check if the parser encountered any errors.
func (p *Parser) Errors() []string {
	return p.errors
}

// peekError is used to add an error to errors when the type of
// peekToken doesn't match the expression.
func (p *Parser) peekError(t token.TokenType) {
	msg := fmt.Sprintf("expected next token to be %s, got %s instead", t, p.peekToken.Type)
	p.errors = append(p.errors, msg)
}

func (p *Parser) noPrefixParseFnError(t token.TokenType) {
	msg := fmt.Sprintf("no prefix parse function for %s found", t)
	p.errors = append(p.errors, msg)
}

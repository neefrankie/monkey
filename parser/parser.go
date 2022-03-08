package parser

import (
	"fmt"
	"monkey/ast"
	"monkey/lexer"
	"monkey/token"
)

// Parser constructs and AST.
// l is a pointer to an instance of the lexer, on which we repeatedly
// call NextToken() to get the next token in the input.
// currToken and peekToken  act exactly like the two pointers our
// lexer has: position and peekPosition.
// Instead of pointing to character in the input, then point to the
// current and the next token.
// We need to look at the curToken, which is the current token under examination,
// to decide  what to do next, and we also need peekToken for this decision
// if curToken doesn't give us enough information.
type Parser struct {
	l         *lexer.Lexer
	curToken  token.Token
	peekToken token.Token
	errors    []string
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{
		l:         l,
		curToken:  token.Token{},
		peekToken: token.Token{},
		errors:    []string{},
	}

	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
	// Construct the root node of AST.
	program := &ast.Program{}
	program.Statements = []ast.Statement{}

	// Iterate over every token in the input until it encounters
	// a token.EOF token.
	for p.curToken.Type != token.EOF {
		stmt := p.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		// Advance both p.curToken and p.peekToken
		p.nextToken()
	}

	// When nothing is left to parse, the *ast.Program root node is returned.
	return program
}

func (p *Parser) parseStatement() ast.Statement {
	switch p.curToken.Type {
	case token.LET:
		return p.parseLetStatement()

	case token.RETURN:
		return p.parseReturnStatement()

	default:
		return nil
	}
}

// parseLetStatement constructs an *ast.LetStatement node
// with the token it's currently sitting on.
func (p *Parser) parseLetStatement() *ast.LetStatement {
	stmt := &ast.LetStatement{
		Token: p.curToken,
	}

	// Expects a token.IDENT token.
	if !p.expectPeek(token.IDENT) {
		return nil
	}

	// Construct an *ast.Identifier node.
	// For x in let x = 5, it should be
	// Token: {Type: token.IDENT, Literal: "x"}
	stmt.Name = &ast.Identifier{
		Token: p.curToken,
		Value: p.curToken.Literal,
	}

	// Expects an equal sign.
	if !p.expectPeek(token.ASSIGN) {
		return nil
	}

	// TODO: we're skipping the expression until we encounter a semicolon
	for !p.curTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

func (p *Parser) parseReturnStatement() *ast.ReturnStatement {
	stmt := &ast.ReturnStatement{
		Token:       p.curToken,
		ReturnValue: nil,
	}

	p.nextToken()

	for !p.curTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

func (p *Parser) curTokenIs(t token.TokenType) bool {
	return p.curToken.Type == t
}

func (p *Parser) peekTokenIs(t token.TokenType) bool {
	return p.peekToken.Type == t
}

// expectPeek is to enforce the correctness of the order of
// tokens by checking the type of the next token.
// Only if the type is correct does it advance the token by
// calling nextToken.
func (p *Parser) expectPeek(t token.TokenType) bool {
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	} else {
		p.peekError(t)
		return false
	}
}

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

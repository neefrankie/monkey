package parser

import (
	"monkey/ast"
	"monkey/token"
)

func (p *Parser) parseStatement() ast.Statement {
	switch p.curToken.Type {
	case token.LET:
		return p.parseLetStatement()

	case token.RETURN:
		return p.parseReturnStatement()

	default:
		return p.parseExpressionStatement()
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

	p.nextToken()

	stmt.Value = p.parseExpression(LOWEST)

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

	stmt.ReturnValue = p.parseExpression(LOWEST)

	for !p.curTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

func (p *Parser) parseExpressionStatement() *ast.ExpressionStatement {
	stmt := &ast.ExpressionStatement{
		Token:      p.curToken,
		Expression: nil,
	}

	stmt.Expression = p.parseExpression(LOWEST)

	if p.peekTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

func (p *Parser) parseBlockStatement() *ast.BlockStatement {
	block := &ast.BlockStatement{
		Token:      p.curToken,
		Statements: nil,
	}

	p.nextToken()

	for !p.curTokenIs(token.RBRACE) && !p.curTokenIs(token.EOF) {
		stmt := p.parseStatement()
		if stmt != nil {
			block.Statements = append(block.Statements, stmt)
		}

		p.nextToken()
	}

	return block
}

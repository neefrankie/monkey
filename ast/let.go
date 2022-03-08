package ast

import "monkey/token"

// LetStatement contains a let statement like let x = 5.
// A let statement consists of:
// * name for the variable;
// * expression on the right side of the equal sign, pointing to any expression
// * the token the AST node is associated with.
type LetStatement struct {
	Token token.Token // the token.LET token
	Name  *Identifier // the x in let x = 5
	Value Expression  // the add(2, 2) * 5 / 10 of a let statement.
}

// satisfy the Statement interface.
func (ls *LetStatement) statementNode() {}

// TokenLiteral implements the Node interface.
func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}

// Identifier holds the identifier in a binding like x
// in let x = 5.
type Identifier struct {
	Token token.Token // the token.IDENT token.
	Value string
}

func (i *Identifier) expressionNode() {}

func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}

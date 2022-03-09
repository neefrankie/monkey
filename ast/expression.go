package ast

import "monkey/token"

// ExpressionStatement is not really a distinct statement;
// it's a statement that consists solely of one expression.
// It can be added the Program.Statements slice.
type ExpressionStatement struct {
	Token      token.Token // The first token of the expression
	Expression Expression
}

func (es *ExpressionStatement) statementNode() {}

func (es *ExpressionStatement) TokenLiteral() string {
	return es.Token.Literal
}

func (es *ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.Expression.String()
	}

	return ""
}

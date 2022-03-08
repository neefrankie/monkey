package ast

// Program is the root node of every AST our parser produces.
// Every valid Monkey program is a series of statements.
// These statements are contained in the Statements.
type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

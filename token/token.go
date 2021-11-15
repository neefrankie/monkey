package token

type TokenType string

// Various toke types.
const (
	// ILLEGAL signifies a token/character we don't know about
	ILLEGAL TokenType = "ILLEGAL"
	// EOF standards for "end of file"
	EOF TokenType = "EOF"

	// IDENT is an identifiers + literals
	IDENT TokenType = "IDENT"
	INT   TokenType = "INT"

	// ASSIGN Operator
	ASSIGN TokenType = "="
	PLUS   TokenType = "+"

	// COMMA Delimiters
	COMMA     TokenType = ","
	SEMICOLON TokenType = ";"

	LPAREN TokenType = "("
	RPAREN TokenType = ")"
	LBRACE TokenType = "{"
	RBRACE TokenType = "}"

	// FUNCTION Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
)

type Token struct {
	Type    TokenType
	Literal string
}

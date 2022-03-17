package token

type TokenType string

// Various toke types.
const (
	// ILLEGAL signifies a token/character we don't know about
	ILLEGAL TokenType = "ILLEGAL"
	// EOF standards for "end of file"
	EOF TokenType = "EOF"

	// ASSIGN Operator
	ASSIGN   TokenType = "="
	PLUS     TokenType = "+"
	MINUS    TokenType = "-"
	BANG     TokenType = "!"
	ASTERISK TokenType = "*"
	SLASH    TokenType = "/"

	LT TokenType = "<"
	GT TokenType = ">"

	EQ     TokenType = "=="
	NOT_EQ TokenType = "!="

	// COMMA Delimiters
	COMMA     TokenType = ","
	SEMICOLON TokenType = ";"

	LPAREN TokenType = "("
	RPAREN TokenType = ")"
	LBRACE TokenType = "{"
	RBRACE TokenType = "}"

	// IDENT is an identifiers + literals
	IDENT TokenType = "IDENT"
	INT   TokenType = "INT"

	// FUNCTION Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	IF       = "IF"
	ELSE     = "ELSE"
	RETURN   = "RETURN"
)

var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
}

// LookupIdent checks if a string is a keyword.
// If it is not one of the specified keywords, then it is
// a variable or function name.
func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}

	return IDENT
}

type Token struct {
	Type    TokenType
	Literal string
}

func NewToken(tokenType TokenType, ch byte) Token {
	return Token{
		Type:    tokenType,
		Literal: string(ch),
	}
}

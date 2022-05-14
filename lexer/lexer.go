package lexer

import "monkey/token"

// Lexer takes source code as input and output the tokens that represent the source code.
// It will go through its input and output the next token it recognizes.
type Lexer struct {
	input        string
	position     int  // current position in input (points to current char)
	readPosition int  // current reading position in input (after current char)
	ch           byte // current char under examination
}

func New(input string) *Lexer {
	l := &Lexer{
		input:        input,
		position:     0,
		readPosition: 0,
		ch:           0,
	}

	l.readChar()

	return l
}

// readChar gives the next character and advance position in the input string.
func (l *Lexer) readChar() {
	// Check whether we have reached the end of the file.
	// If there's the case it sets l.ch to 0, which the ASCII code for the "NUL" character
	// and signifies either we haven't read anything yet, or end of file.
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	// Point to current char under examination.
	l.position = l.readPosition
	// Move to next char.
	l.readPosition += 1
}

// peekChar is similar to readChar, except that it doesn't
// move position and readPosition.
// We only want to peek ahead in the input and not move
// around in it, so we know what a call to readChar would
// return.
func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}

// NextToken looks at the current character under examination and return a token
// depending on which character it is.
// Before returning the token we advance our pointers into the input
// so when we call NextToken() again the l.ch field is already updated.
func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	switch l.ch {
	case '=':
		// Look ahead in the input and then determine
		// whether to return a token for = or ==.
		if l.peekChar() == '=' {
			ch := l.ch   // Save char at current position
			l.readChar() // Move to next position
			tok = token.Token{
				Type:    token.EQ,
				Literal: string(ch) + string(l.ch),
			}
		} else {
			tok = token.NewToken(token.ASSIGN, l.ch)
		}

	case '+':
		tok = token.NewToken(token.PLUS, l.ch)

	case '-':
		tok = token.NewToken(token.MINUS, l.ch)

	case '!':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = token.Token{
				Type:    token.NOT_EQ,
				Literal: string(ch) + string(l.ch),
			}
		} else {
			tok = token.NewToken(token.BANG, l.ch)
		}

	case '/':
		tok = token.NewToken(token.SLASH, l.ch)

	case '*':
		tok = token.NewToken(token.ASTERISK, l.ch)

	case '<':
		tok = token.NewToken(token.LT, l.ch)

	case '>':
		tok = token.NewToken(token.GT, l.ch)

	case ';':
		tok = token.NewToken(token.SEMICOLON, l.ch)

	case ',':
		tok = token.NewToken(token.COMMA, l.ch)

	case '(':
		tok = token.NewToken(token.LPAREN, l.ch)

	case ')':
		tok = token.NewToken(token.RPAREN, l.ch)

	case '{':
		tok = token.NewToken(token.LBRACE, l.ch)

	case '}':
		tok = token.NewToken(token.RBRACE, l.ch)

	case '[':
		tok = token.NewToken(token.LBRACKET, l.ch)

	case ']':
		tok = token.NewToken(token.RBRACKET, l.ch)

	case 0:
		tok.Literal = ""
		tok.Type = token.EOF

	case '"':
		tok.Type = token.STRING
		tok.Literal = l.readString()

	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		} else if isDigit(l.ch) {
			tok.Type = token.INT
			tok.Literal = l.readNumber()
			return tok
		} else {
			tok = token.NewToken(token.ILLEGAL, l.ch)
		}
	}

	l.readChar()
	return tok
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}

	// Now l.position points to one char after the identifier.
	return l.input[position:l.position]
}

func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) readString() string {
	position := l.position + 1
	for {
		l.readChar()
		if l.ch == '"' || l.ch == 0 {
			break
		}
	}

	return l.input[position:l.position]
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

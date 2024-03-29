package lexer

import "jb-lang/token"

type Lexer struct {
	input        string
	position     int  // current position in input (points to current char)
	readPosition int  // current reading position in input (after current char)
	ch           byte // current char under examination
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	switch l.ch {
	case '=':
		tok = newToken(token.ASSIGN, l.ch)
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '{':
		tok = newToken(token.LEFT_BRACE, l.ch)
	case '}':
		tok = newToken(token.RIGHT_BRACE, l.ch)
	case '(':
		tok = newToken(token.LEFT_PAREN, l.ch)
	case ')':
		tok = newToken(token.RIGHT_PAREN, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case ':':
		tok = newToken(token.COLON, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF

	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		} else if isSymbol(l.ch) {
			tok.Literal = l.readOperator()
			tok.Type = token.LookupOperator(tok.Literal)
			return tok
		} else if isDigit(l.ch) {
			tok.Literal = l.readNumber()
			tok.Type = token.INT
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}
	}

	l.readChar()
	return tok
}

func (l *Lexer) readIdentifier() string {
	start := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	end := l.position
	return l.input[start:end]
}

func (l *Lexer) readOperator() string {

	start := l.position
	for isSymbol(l.ch) {
		l.readChar()
	}
	end := l.position
	return l.input[start:end]
}

func (l *Lexer) readNumber() string {
	start := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	end := l.position
	return l.input[start:end]
}

func isSymbol(x byte) bool {
	return x == '-' || x == '>'
}

func isDigit(x byte) bool {
	return '0' <= x && x <= '9'
}

func isLetter(x byte) bool {
	return 'a' <= x && x <= 'z' || 'A' <= x && x <= 'Z' || x == '_'
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

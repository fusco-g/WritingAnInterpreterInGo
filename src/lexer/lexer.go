package lexer

import (
	"interpeter/token"
)

type Lexer struct {
	input                string
	previousCharPosition int
	readCharPosition     int
	charReaded           byte
}

func NewLexer(input string) *Lexer {
	return &Lexer{
		input:                input,
		previousCharPosition: 0,
		readCharPosition:     0,
		charReaded:           input[0],
	}
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	switch l.charReaded {
	case '=':
		// complete the token
		if l.peekChar() == '=' {
			ch := l.charReaded
			l.readChar()
			literal := string(ch) + string(l.charReaded)
			// produce the token
			tok = token.Token{
				Type:    token.EQUAL,
				Literal: literal,
			}
		} else {
			tok = newToken(token.ASSIGN, l.charReaded)
		}
	case '+':
		tok = newToken(token.PLUS, l.charReaded)
	case '-':
		tok = newToken(token.MIN, l.charReaded)
	case '!':
		// complete the token
		if l.peekChar() == '=' {
			ch := l.charReaded
			l.readChar()
			literal := string(ch) + string(l.charReaded)
			// produce the token
			tok = token.Token{
				Type:    token.NOT_EQUAL,
				Literal: literal,
			}
		} else {
			tok = newToken(token.BANG, l.charReaded)
		}
	case '/':
		tok = newToken(token.DIV, l.charReaded)
	case '*':
		tok = newToken(token.MUL, l.charReaded)
	case '<':
		tok = newToken(token.MINOR, l.charReaded)
	case '>':
		tok = newToken(token.MAJOR, l.charReaded)
	case ';':
		tok = newToken(token.SEMICOLON, l.charReaded)
	case ',':
		tok = newToken(token.COMMA, l.charReaded)
	case '{':
		tok = newToken(token.O_BRACE, l.charReaded)
	case '}':
		tok = newToken(token.C_BRACE, l.charReaded)
	case '(':
		tok = newToken(token.O_ROUND_BRACKET, l.charReaded)
	case ')':
		tok = newToken(token.C_ROUND_BRACKET, l.charReaded)
	case 0:
		// if the readed character is 0 it means
		// that the source code is ended
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		// check if the character readed is
		// a identifier or a number
		if isLetter(l.charReaded) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.SearchKeyword(tok.Literal)
			return tok
		} else if isDigit(l.charReaded) {
			tok.Type = token.INT
			tok.Literal = l.readNumber()
		} else {
			tok = newToken(token.ILLEGAL, l.charReaded)
		}
	}

	l.readChar()
	return tok
}

func (l *Lexer) skipWhitespace() {
	for l.charReaded == ' ' || l.charReaded == '\t' || l.charReaded == '\n' || l.charReaded == '\r' {
		l.readChar()
	}
}

func (l *Lexer) readChar() {
	if l.readCharPosition >= len(l.input) {
		// max reached
		l.charReaded = 0
	} else {
		l.charReaded = l.input[l.readCharPosition]
	}
	l.previousCharPosition = l.readCharPosition
	l.readCharPosition += 1
}

func (l *Lexer) peekChar() byte {
	if l.readCharPosition >= len(l.input) {
		return 0
	}
	// return the next character. = ---> ==
	return l.input[l.readCharPosition]
}

func (l *Lexer) readIdentifier() string {
	position := l.previousCharPosition
	// read the input until is a character
	// so is possible to increment
	// readCharPosition and previousCharPosition
	for isLetter(l.charReaded) {
		l.readChar()
	}
	// return the identifier by slincing the
	// input string
	return l.input[position:l.previousCharPosition]
}

func (l *Lexer) readNumber() string {
	position := l.previousCharPosition
	// read the input until is a digit
	// so i possible to increment
	// readCharPosition and previousCharPosition
	for isDigit(l.charReaded) {
		l.readChar()
	}
	// return the identified by sliciding the
	// input string
	return l.input[position:l.previousCharPosition]
}

func newToken(tokenType token.TokenType, char byte) token.Token {
	return token.Token{
		Type:    tokenType,
		Literal: string(char),
	}
}

func isLetter(char byte) bool {
	switch {
	case char >= 'a' && char <= 'z', char >= 'A' && char <= 'Z', char == '_':
		return true
	}
	return false
}

func isDigit(char byte) bool {
	switch {
	case char >= '0' && char <= '9':
		return true
	}
	return false
}

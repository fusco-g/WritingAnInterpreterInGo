package parser

import "interpeter/token"

type Parser struct {
	tokenInstance token.Token

	// pointer to the current token
	currentToken token.Token

	// pointer to the next token
	nextToken token.Token
}

func NewParser() *Parser {
	return &Parser{
		// TODO
	}
}

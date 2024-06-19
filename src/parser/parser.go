package parser

import (
	"interpeter/ast"
	"interpeter/lexer"
	"interpeter/token"
)

type Parser struct {
	lexerInstance lexer.Lexer

	// pointer to the current token
	currentToken token.Token

	// pointer to the next token
	nextToken token.Token
}

func NewParser(lexer lexer.Lexer) *Parser {
	p := &Parser{lexerInstance: lexer}
	p.nToken()
	p.nToken()
	return p
}

func (p *Parser) nToken() {
	p.currentToken = p.nextToken
	p.nextToken = p.lexerInstance.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
	return nil
}

package ast

import "go/token"

type Node interface {
	TokenLiteral() string
}

type Statement interface {
	// type embedding
	// hinerit Node interface methods
	Node
	statementNode()
}

type Expression interface {
	// type embedding
	// hinerit Node interface methods
	Node
	expressionNode()
}

// contain all the input statements
type Program struct {
	Statements []Statement
}

// produce the root node
func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

type LetStatement struct {
	letToken token.Token
	// TODO
}

package ast

import "interpeter/token"

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
// witch are the nodes of the AST
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

type Identifier struct {
	Token token.Token // IDENTIFIER token
	Value string      // let x = 5, x is the identifier in the statement
}

// LET statement structure
// let x = 6 * 6
// letToken = let
// Name  = x
// Value = 6 * 6
type LetStatement struct {
	letToken token.Token
	Name     *Identifier
	Value    Expression
}

// LetStatement implements Statement interface
func (l *LetStatement) statementNode() {

}

// implementing Node interface
func (l *LetStatement) TokenLiteral() string {
	return l.letToken.Literal
}

// implementing Expression interface
func (i *Identifier) expressionNode() {

}

// implementing Node interface
func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}

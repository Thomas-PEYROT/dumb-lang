package ast

import "bytes"

type Node interface {
	TokenLiteral() string // Returns the token literal value
	String() string
}

// Statement is a type of Node. It's an instruction that doesn't "return" anything, by opposition to Expression.
type Statement interface {
	Node
	statementNode()
}

// Expression is a type of Node. It's an instruction that "returns" a value, by opposition to Statement.
type Expression interface {
	Node
	expressionNode()
}

// Program is basically a list of Statement.
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

func (p *Program) String() string {
	var out bytes.Buffer
	for _, s := range p.Statements {
		out.WriteString(s.String())
	}
	return out.String()
}

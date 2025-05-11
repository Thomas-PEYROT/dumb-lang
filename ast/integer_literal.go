package ast

import "dumb-lang/token"

type IntegerLiteral struct {
	Token token.Token
	Value int64
}

func (il *IntegerLiteral) expressionNode()      {}                          // Implementing Expression interface
func (il *IntegerLiteral) TokenLiteral() string { return il.Token.Literal } // Implementing Node interface
func (il *IntegerLiteral) String() string       { return il.Token.Literal } // Implementing Node interface

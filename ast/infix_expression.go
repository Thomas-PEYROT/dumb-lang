package ast

import (
	"bytes"
	"dumb-lang/token"
)

type InfixExpression struct {
	Token    token.Token // The operator token, e.g. +
	Left     Expression
	Operator string
	Right    Expression
}

func (ie *InfixExpression) expressionNode()      {}                          // Implementing Expression interface
func (ie *InfixExpression) TokenLiteral() string { return ie.Token.Literal } // Implementing Node interface
func (ie *InfixExpression) String() string {
	var out bytes.Buffer
	out.WriteString("(")
	out.WriteString(ie.Left.String())
	out.WriteString(" " + ie.Operator + " ")
	out.WriteString(ie.Right.String())
	out.WriteString(")")
	return out.String()
} // Implementing Node interface

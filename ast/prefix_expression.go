package ast

import (
	"bytes"
	"dumb-lang/token"
)

type PrefixExpression struct {
	Token    token.Token // The prefix token, e.g.
	Operator string
	Right    Expression
}

func (pe *PrefixExpression) expressionNode()      {} // Implement Expression interface
func (pe *PrefixExpression) TokenLiteral() string { return pe.Token.Literal }
func (pe *PrefixExpression) String() string {
	var out bytes.Buffer
	out.WriteString("(")
	out.WriteString(pe.Operator)
	out.WriteString(pe.Right.String())
	out.WriteString(")")
	return out.String()
}

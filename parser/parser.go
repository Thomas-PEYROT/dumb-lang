package parser

import (
	"dumb-lang/ast"
	"dumb-lang/lexer"
	"dumb-lang/token"
)

type Parser struct {
	l            *lexer.Lexer
	currentToken token.Token // Current token
	peekToken    token.Token // Next token
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}

	// Read two tokens, so curToken and peekToken are both set
	p.nextToken()
	p.nextToken()
	return p
}

func (p *Parser) nextToken() {
	p.currentToken = p.peekToken
	p.peekToken = p.l.NextToken()
}
func (p *Parser) ParseProgram() *ast.Program {
	return nil
}

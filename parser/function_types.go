package parser

import "dumb-lang/ast"

type prefixParseFn func() ast.Expression
type infixParseFn func(ast.Expression) ast.Expression

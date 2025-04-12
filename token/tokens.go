package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Identifiers + literals
	IDENT = "IDENT"
	INT   = "INT"

	// Operators
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	BANG     = "!"
	ASTERISK = "*"
	SLASH    = "/"
	LT       = "<"
	GT       = ">"

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"

	// Function
	FUNCTION = "FUNCTION"
	LET      = "LET"
)

// Used to differentiate reserved keywords and classic identifiers
var keywords = map[string]TokenType{
	"fn":  FUNCTION,
	"let": LET,
}

// LookupIdent : Util function to check whether an identifier is a reserved keyword or not
func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}

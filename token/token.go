package token

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// IDENTIFIERS
	IDENT = "IDENT"

	// LITERALS
	INT  = "INT"
	BOOL = "BOOL"

	// OPERATORS
	ASSIGN = "="
	PLUS   = "+"

	// DELIMITERS
	COMMA      = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// KEYWORDS
	FUNCTION = "fn"
	VAR      = "var"
)

type Token struct {
	Type    string
	Literal string
}

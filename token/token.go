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
	SEMILCOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// KEYWORDS
	FUNCTION = "fn"
	VAR      = "var"
)

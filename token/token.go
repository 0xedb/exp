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
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// KEYWORDS
	FUNCTION = "fn"
	VAR      = "var"
)

var kewywords = map[string]string{
	"fn":  "fn",
	"var": "var",
}

func LookUpIdentifier(id string) string {
	if tok, ok := kewywords[id]; ok {
		return tok
	}

	return IDENT
}

type Token struct {
	Type    string
	Literal string
}

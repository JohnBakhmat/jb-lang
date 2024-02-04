package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	IDENT = "IDENT"
	INT   = "INT"

	ASSIGN    = "="
	PLUS      = "+"
	COMMA     = ","
	SEMICOLON = ";"
	COLON     = ":"

	LEFT_PAREN  = "("
	RIGHT_PAREN = ")"
	LEFT_BRACE  = "{"
	RIGHT_BRACE = "}"

	RIGHT_ARROW = "->"

	FUNCTION = "FUNCTION"
	VAR      = "VAR"

	TYPE_INT = "TYPE_INT"
)

var keywords = map[string]TokenType{
	"fn":  FUNCTION,
	"var": VAR,
	"int": TYPE_INT,
}

var operators = map[string]TokenType{
	"->": RIGHT_ARROW,
}

func LookupOperator(op string) TokenType {
	if tok, ok := operators[op]; ok {
		return tok
	}
	return ILLEGAL
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}

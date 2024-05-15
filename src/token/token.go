package token

type (
	TokenType string

	Token struct {
		// here goes the value of the constants
		Type TokenType

		// value of the token
		Literal string
	}
)

const (
	// special characters
	ILLEGAL, EOF = "ILLEGAL", "EOF"

	// arithmetic logic operators
	PLUS       = "+"
	MIN        = "-"
	DIV        = "/"
	MUL        = "*"
	EQUALS     = "=="
	NOT_EQUALS = "!="
	MINOR      = "<"
	MAJOR      = ">"
	ASSIGN     = "="

	// delimiters
	COMMA, SEMICOLO = ",", ";"

	// parentheses
	O_ROUND_BRACKET = "("
	C_ROUND_BRACKET = ")"
	O_BRACE         = "{"
	C_BRACE         = "}"

	// keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
	IF       = "IF"
	ELSE     = "ELSE"
	RETURN   = "RETURN"

	// types
	INT = "INT"

	// identifiers
	IDENTIFIER = "IDENTIFIER"

	// boolean literals
	TRUE  = "TRUE"
	FALSE = "FALSE"
)

var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
}

func searchKeyword(keywordToSearch string) TokenType {
	// comma ok pattern. ok contain a boolean value
	// if keyword exist return token
	if token, ok := keywords[keywordToSearch]; ok {
		return token
	}
	return IDENTIFIER
}

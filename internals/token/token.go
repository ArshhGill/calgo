package token

const (
	PLUS     = "+"
	MINUS    = "-"
	MULTIPLY = "*"
	DIVIDE   = "/"
	MODULO   = "%"
	EXPONENT = "**"
	INT      = "int"
	NULL     = "null"
	ILLEGAL  = "cocaine"
	EOL      = "eol"
)

type Token struct {
	Type    string
	Literal string
}

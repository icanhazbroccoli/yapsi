package token

import "strings"

type TokenType string

const (
	IDENT     TokenType = "IDENT"
	NUMBER              = "NUMBER"
	STRING              = "STRING"
	POINTER             = "POINTER"
	CHAR                = "CHAR"
	BOOL                = "BOOL"
	EOF                 = "EOF"
	ILLEGAL             = "ILLEGAL"
	PLUS                = "+"
	MINUS               = "-"
	ASTERISK            = "*"
	SLASH               = "/"
	EQUAL               = "="
	LESS                = "<"
	GREATER             = ">"
	LBRACKET            = "["
	RBRACKET            = "]"
	PERIOD              = "."
	COMMA               = ","
	COLON               = ":"
	SEMICOLON           = ";"
	PTR                 = "^"
	LPAREN              = "("
	RPAREN              = ")"
	NOTEQUAL            = "<>"
	LTEQL               = "<="
	GTEQL               = ">="
	NAMED               = ":="
	DOTDOT              = ".."
	AND                 = "AND"
	ARRAY               = "ARRAY"
	BEGIN               = "BEGIN"
	CASE                = "CASE"
	CONST               = "CONST"
	DIV                 = "DIV"
	DO                  = "DO"
	DOWNTO              = "DOWNTO"
	ELSE                = "ELSE"
	END                 = "END"
	FALSE               = "FALSE"
	FILE                = "FILE"
	FOR                 = "FOR"
	FUNCTION            = "FUNCTION"
	GOTO                = "GOTO"
	IF                  = "IF"
	IN                  = "IN"
	LABEL               = "LABEL"
	MOD                 = "MOD"
	NIL                 = "NIL"
	NOT                 = "NOT"
	OF                  = "OF"
	OR                  = "OR"
	PACKED              = "PACKED"
	PROCEDURE           = "PROCEDURE"
	PROGRAM             = "PROGRAM"
	RECORD              = "RECORD"
	REPEAT              = "REPEAT"
	SET                 = "SET"
	THEN                = "THEN"
	TO                  = "TO"
	TRUE                = "TRUE"
	TYPE                = "TYPE"
	UNTIL               = "UNTIL"
	VAR                 = "VAR"
	WHILE               = "WHILE"
	WITH                = "WITH"
)

type Token struct {
	Type    TokenType
	Literal string
}

var keywords = map[string]TokenType{
	"and":       AND,
	"begin":     BEGIN,
	"case":      CASE,
	"const":     CONST,
	"div":       DIV,
	"do":        DO,
	"downto":    DOWNTO,
	"else":      ELSE,
	"end":       END,
	"false":     FALSE,
	"file":      FILE,
	"for":       FOR,
	"function":  FUNCTION,
	"goto":      GOTO,
	"if":        IF,
	"label":     LABEL,
	"mod":       MOD,
	"nil":       NIL,
	"not":       NOT,
	"of":        OF,
	"or":        OR,
	"packed":    PACKED,
	"procedure": PROCEDURE,
	"program":   PROGRAM,
	"repeat":    REPEAT,
	"then":      THEN,
	"to":        TO,
	"true":      TRUE,
	"type":      TYPE,
	"until":     UNTIL,
	"var":       VAR,
	"while":     WHILE,
	"with":      WITH,
}

func LookupIdent(ident string) TokenType {
	ident = strings.ToLower(ident)
	if t, ok := keywords[ident]; ok {
		return t
	}
	return IDENT
}

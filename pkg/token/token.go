package token

import "strings"

const (
	IDENT   = "IDENT"
	EOF     = "EOF"
	ILLEGAL = "ILLEGAL"

	BOOLEAN     = "boolean"
	BYTE        = "byte"
	CARDINAL    = "cardinal"
	CHAR        = "char"
	CURRENCY    = "currency"
	EXTENDED    = "extended"
	INT64       = "int64"
	INTEGER     = "integer"
	LONGINT     = "longint"
	NUMBER      = "number"
	POINTER     = "pointer"
	REAL        = "real"
	SHORTINT    = "shortint"
	SMALLINT    = "smallint"
	WORD        = "word"
	ARRAY       = "array"
	CLASS       = "class"
	OBJECT      = "object"
	RECORD      = "record"
	SET         = "set"
	STRING      = "string"
	SHORTSTRING = "shortstring"

	AND       = "and"
	BEGIN     = "begin"
	CASE      = "case"
	CONST     = "const"
	DIV       = "div"
	DO        = "do"
	DOWNTO    = "downto"
	ELSE      = "else"
	END       = "end"
	FILE      = "file"
	FOR       = "for"
	FUNCTION  = "function"
	GOTO      = "goto"
	IF        = "if"
	LABEL     = "label"
	MOD       = "mod"
	NOT       = "not"
	OF        = "of"
	OR        = "or"
	PACKED    = "packed"
	PROCEDURE = "procedure"
	PROGRAM   = "program"
	REPEAT    = "repeat"
	THEN      = "then"
	TO        = "to"
	TYPE      = "type"
	UNTIL     = "until"
	USES      = "uses"
	VAR       = "var"
	WHILE     = "while"
	WITH      = "with"

	PLUS      = "+"
	MINUS     = "-"
	ASTERISK  = "*"
	SLASH     = "/"
	EQUAL     = "="
	LESS      = "<"
	GREATER   = ">"
	LBRACKET  = "["
	RBRACKET  = "]"
	PERIOD    = "."
	COMMA     = ","
	COLON     = ":"
	SEMICOLON = ";"
	PTR       = "^"
	LPAREN    = "("
	RPAREN    = ")"
	NOTEQUAL  = "<>"
	LTEQL     = "<="
	GTEQL     = ">="
	NAMED     = ":="
	DOTDOT    = ".."

	NIL   = "nil"
	TRUE  = "true"
	FALSE = "false"
)

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

var keywords = map[string]TokenType{
	"boolean":     BOOLEAN,
	"byte":        BYTE,
	"cardinal":    CARDINAL,
	"char":        CHAR,
	"currency":    CURRENCY,
	"extended":    EXTENDED,
	"int64":       INT64,
	"integer":     INTEGER,
	"longint":     LONGINT,
	"pointer":     POINTER,
	"real":        REAL,
	"shortint":    SHORTINT,
	"smallint":    SMALLINT,
	"word":        WORD,
	"array":       ARRAY,
	"class":       CLASS,
	"object":      OBJECT,
	"record":      RECORD,
	"set":         SET,
	"string":      STRING,
	"shortstring": SHORTSTRING,

	"and":       AND,
	"begin":     BEGIN,
	"case":      CASE,
	"const":     CONST,
	"div":       DIV,
	"do":        DO,
	"downto":    DOWNTO,
	"else":      ELSE,
	"end":       END,
	"file":      FILE,
	"for":       FOR,
	"function":  FUNCTION,
	"goto":      GOTO,
	"if":        IF,
	"label":     LABEL,
	"mod":       MOD,
	"not":       NOT,
	"of":        OF,
	"or":        OR,
	"packed":    PACKED,
	"procedure": PROCEDURE,
	"program":   PROGRAM,
	"repeat":    REPEAT,
	"then":      THEN,
	"to":        TO,
	"type":      TYPE,
	"until":     UNTIL,
	"uses":      USES,
	"var":       VAR,
	"while":     WHILE,
	"with":      WITH,

	"nil":   NIL,
	"true":  TRUE,
	"false": FALSE,
}

func LookupIdent(ident string) TokenType {
	ident = strings.ToLower(ident)
	if t, ok := keywords[ident]; ok {
		return t
	}
	return IDENT
}

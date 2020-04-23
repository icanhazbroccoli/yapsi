package repl

import (
	"bufio"
	"io"
	"strings"

	"yapsi/pkg/lexer"
	"yapsi/pkg/parser"
)

const (
	PROMPT = ">> "
)

func Run(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		io.WriteString(out, PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		lex := lexer.New(strings.NewReader(line))
		p := parser.New(lex)
		program := p.ParseProgram()
		if err := p.Error(); err != nil {
			io.WriteString(out, "ERROR: "+err.Error())
			continue
		}
		io.WriteString(out, program.String())

		io.WriteString(out, "\n")
	}
}

package repl

import (
	"bufio"
	"fmt"
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
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		lex, err := lexer.New(strings.NewReader(line))
		if err != nil {
			io.WriteString(out, fmt.Sprintf("ERROR: %s", err))
			continue
		}
		p, err := parser.New(lex)
		if err != nil {
			io.WriteString(out, fmt.Sprintf("ERROR: %s", err))
			continue
		}
		program, err := p.ParseProgram()
		if err != nil {
			io.WriteString(out, fmt.Sprintf("ERROR: %s", err))
			continue
		}
		io.WriteString(out, program.String())

		io.WriteString(out, "\n")
	}
}

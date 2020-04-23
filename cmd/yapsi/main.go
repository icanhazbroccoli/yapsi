package main

import (
	"bufio"
	"fmt"
	"os"

	"yapsi/pkg/lexer"
	"yapsi/pkg/parser"
)

func main() {
	filename := os.Args[1]
	file, err := os.Open(filename)
	if err != nil {
		panic(err.Error())
	}
	lex := lexer.New(bufio.NewReader(file))
	p := parser.New(lex)
	program := p.ParseProgram()
	if err := p.Error(); err != nil {
		panic(err.Error())
	}
	fmt.Println(program.String())
}

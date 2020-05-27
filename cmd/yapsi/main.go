package main

import (
	"bufio"
	"fmt"
	"os"

	"yapsi/pkg/lexer"
	"yapsi/pkg/parser"
	"yapsi/pkg/printer"
)

func main() {
	filename := os.Args[1]
	file, err := os.Open(filename)
	if err != nil {
		panic(err.Error())
	}
	lex := lexer.New(bufio.NewReader(file))
	p := parser.New(lex)
	program, err := p.ParseProgram()
	if err != nil {
		panic(err.Error())
	}
	printer := &printer.AstPrinter{}
	fmt.Println(program.Visit(printer).(string))
}

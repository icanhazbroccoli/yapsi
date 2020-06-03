package main

import (
	"bufio"
	"fmt"
	"os"

	"yapsi/pkg/interpreter"
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
	str, err := program.Visit(printer)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(str.(string))

	var it8 interpreter.Interpreter
	if err := it8.Evaluate(program); err != nil {
		panic(err.Error())
	}
}

package main

import (
	"os"

	"yapsi/pkg/repl"
)

func main() {
	repl.Run(os.Stdin, os.Stdout)
}

package main

import (
	"fmt"
	"io"
	"os"

	"yapsi/pkg/repl"
)

func main() {
	repl.Run(os.Stdin, os.Stdout)
}

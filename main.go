package main

import (
	"magnes/repl"
	"os"
)

func main() {
	repl.StartParser(os.Stdin, os.Stdout)
}

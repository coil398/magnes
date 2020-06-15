package main

import (
	"magnes/repl"
	"os"
)

func main() {
	repl.Start(os.Stdin, os.Stdout)
}

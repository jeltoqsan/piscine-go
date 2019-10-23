package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	for _, r := range os.Args[0] {
		z01.PrintRune(r)
	}
	z01.PrintRune(10)
}

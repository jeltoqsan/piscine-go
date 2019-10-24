package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	counter := 0
	for i := range os.Args {
		counter = i
	}
	for i := counter; i > 0; i = i - 1 {
		for _, r := range os.Args[i] {
			z01.PrintRune(r)
		}
		z01.PrintRune(10)
	}
}

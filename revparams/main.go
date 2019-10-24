package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	x := 0
	for y := range os.Args {
		x = y
	}
	for y := x; y > 0; y = y - 1 {
		for _, z := range os.Args[y] {
			z01.PrintRune(z)
		}
		z01.PrintRune(10)
	}
}

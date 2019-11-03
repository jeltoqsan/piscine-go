package main

import "github.com/01-edu/z01"

func main() {
	var x rune = 97
	for i := x; i <= 122; i++ {
		z01.PrintRune(i)
	}
	z01.PrintRune(10)
}

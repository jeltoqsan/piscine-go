package main

import "github.com/01-edu/z01"

func main() {
	for latin := 'a'; latin <= 'z'; latin++ {
		z01.PrintRune(latin)
	}
	z01.PrintRune('\n')
} 

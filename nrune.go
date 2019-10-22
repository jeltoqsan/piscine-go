package main

//package piscine

import "github.com/01-edu/z01"

func NRune(s string, n int) rune {
	x := 0
	for _, z := range s {
		x = x + 1
		if x == n {
			return z
		}
	}
	return 0
}

func main() {
	z01.PrintRune(NRune("Hello!", 3))
	z01.PrintRune(NRune("Salut!", 2))
	z01.PrintRune(NRune("Ola!", 4))
	z01.PrintRune('\n')
}

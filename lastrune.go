// package piscine
package main

import "github.com/01-edu/z01"

	func FirstRune(s string) rune {
		var a rune
		for _, z := range s {
			a = z
		}
		return a
	}
}
func main() {
	z01.PrintRune(LastRune("Hello!"))
	z01.PrintRune(LastRune("Salut!"))
	z01.PrintRune(LastRune("Ola!"))
	z01.PrintRune(10)
}

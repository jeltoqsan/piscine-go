package piscine

import "github.com/01-edu/z01"

func PrintWordsTables(table []string) {
	for _, r := range table {
		for _, char := range []rune(r) {
			z01.PrintRune(char)
		}

		z01.PrintRune('\n')
	}
}

package main

import (
	"github.com/01-edu/z01"
)

func PrintNbrBase(nbr int, base string) {
	counter := 0
	for _, r1 := range base {
		flag := 0
		if r1 == '+' || r1 == '-' {
			z01.PrintRune('N')
			z01.PrintRune('V')
			return
		}
		for _, r2 := range base {
			if r1 == r2 {
				flag++
			}
			if flag == 2 {
				z01.PrintRune('N')
				z01.PrintRune('V')
				return
			}
		}
		counter++
	}
	if counter < 2 {
		z01.PrintRune('N')
		z01.PrintRune('V')
		return
	}
	res := ""
	ra := []rune(base)
	if nbr < 0 {
		z01.PrintRune('-')
	}
	for ; nbr != 0; nbr /= counter {
		index := nbr % counter
		if index < 0 {
			index = -index
		}
		res = string(ra[index]) + res
	}
	for _, r := range res {
		z01.PrintRune(r)
	}
}

func main() {
	PrintNbrBase(125, "0123456789")
	z01.PrintRune('\n')
	PrintNbrBase(-125, "01")
	z01.PrintRune('\n')
	PrintNbrBase(125, "0123456789ABCDEF")
	z01.PrintRune('\n')
	PrintNbrBase(-125, "choumi")
	z01.PrintRune('\n')
	PrintNbrBase(125, "aa")
	z01.PrintRune('\n')
}

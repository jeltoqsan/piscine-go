package piscine

import "github.com/01-edu/z01"

func PrintNbr(n int)
	z := 0
	if n == 0 {
		z01.PrintRune(z)
		return
	}
	for j := 1; j <= n%10; j++ {
		z++
	}
	for j := -1; j >= n%10; j-- {
		z++
	}
	if n/10 != 0 {
		PrintNbr(n / 10)
	}
	z01.PrintRune(z)
	return
}
func PrintNbr(x int) {
	if x < 0 {
		z01.PrintRune('-')
	}
	PrintNbr(x)
}

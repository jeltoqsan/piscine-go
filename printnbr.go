package piscine

import "github.com/01-edu/z01"

func PrintNbr(n int)

	w := '0'
	if n == 0 {
		z01.PrintRune(w)
		return
	}
	for j := 1; j <= n%10; j++ {
		w++
	}
	for j := -1; j >= n%10; j-- {
		w++
	}
	if n/10 != 0 {
		PrintNbr(n / 10)
	}
	z01.PrintRune(w)
	return
}
func PrintNbr(x int) {
	if x < 0 {
		z01.PrintRune('-')
	}
	PrintNbr(x)


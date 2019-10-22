package piscine

import "github.com/01-edu/z01"

func PrintNbr(n int) {
	i := '0'
	if n == 0 {
		z01.PrintRune('0')
		return
	}
	for j := 1; j <= n%10; j++ {
		i++
	}
	for j := -1; j >= n%10; j-- {
		i++
	}
	if n/10 != 0 {
		PrintRune(n / 10)
	}
	z01.PrintRune(i)
	return
}

func PrintRune(n int) {

	if n < 0 {
		z01.PrintRune('-')
	}
	PrintRune(n)
}

package piscine

import "github.com/01-edu/z01"

func PrintNbr(n int) {
	var i int //i := '0'
	i = 0
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
		PrintNbr(n / 10)
	}
	z01.PrintRune(i)
	return
}

func PrintNbr(n int) {

	if n < 0 {
		z01.PrintRune('-')
	}
	PrintNum(n)
}

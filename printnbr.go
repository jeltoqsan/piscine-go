//package main

package piscine

import "github.com/01-edu/z01"

func PrintNb(nb int) {
	i := '0'
	if nb == 0 {
		z01.PrintRune('0')
		return
	}
	for j := 1; j <= nb%10; j++ {
		i++
	}
	for j := -1; j >= nb%10; j-- {
		i++
	}
	if nb/10 != 0 {
		PrintNb(nb / 10)
	}
	z01.PrintRune(i)
	return
}

func PrintNbr(n int) {

	if n < 0 {
		z01.PrintRune('-')
	}
	PrintNb(n)
}

/*func main() {
	PrintNbr(-123)
	PrintNbr(0)
	PrintNbr(123)
	z01.PrintRune(10)
}*/

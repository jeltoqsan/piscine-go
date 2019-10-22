package piscine

//package main

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	i := '0'
	if n == 0 {
		z01.PrintRune('0')
		return
	}
	for j := 1; j <= n%10; j++ {
		i++
	}
	if n/10 != 0 {
		PrintNbrInOrder(n / 10)
	}
	z01.PrintRune(i)
	return
}

/*func main() {
	PrintNbrInOrder(321)
	PrintNbrInOrder(0)
	PrintNbrInOrder(321)
	z01.PrintRune(10)
}*/

package main

import "github.com/01-edu/z01"

func main() {
	var a rune = 48
	for i := a; i <= 57; i = i + 1 { // for - цикл, выполняется усл.выражение если 0 меньше или равно 9, то выводится 0, после: присвоить i+1 затем 1, 2 итд. Как только 9=9 - цикл останавливается
		z01.PrintRune(i) // ascii code
	}
	z01.PrintRune(10)
}

package piscine

import "github.com/01-edu/z01"

/*package main

import "github.com/01-edu/z01"*/

func Raid1a(x, y int) {
	for i := 1; i <= y; i++ {
		for j := 1; j <= x; j++ {
			if i == 1 && j == 1 || i == 1 && j == x || i == y && j == 1 || i == y && j == x {
				z01.PrintRune('o')
			} else if i < y && i > 1 && j == 1 || i < y && i > 1 && j == x {
				z01.PrintRune('|')
			} else if i == 1 && j > 1 && j < x || i == y && j > 1 && j < x {
				z01.PrintRune('-')
			} else {
				z01.PrintRune(' ')
			}
		}
		if x > 0 && y > 0 {
			z01.PrintRune(10)
		}
	}
}

/*func main() {
	Raid1a(5, 3)
}*/

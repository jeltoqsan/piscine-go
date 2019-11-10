package piscine

/*package main

import . "github.com/01-edu/z01"*/

func Raid1b(x, y int) {
	for i := 1; i <= y; i++ {
		for j := 1; j <= x; j++ {
			if i == 1 && j == 1 {
				PrintRune('/')
			} else if i == 1 && j == x {
				PrintRune(92)
			} else if i == y && j == 1 {
				PrintRune(92)
			} else if i == y && j == x {
				PrintRune('/')
			} else if i == 1 && j > 1 && j < x || i == y && j > 1 && j < x || i > 1 && i < y && j == 1 || i > 1 && i < y && j == x {
				PrintRune('*')
			} else {
				PrintRune(' ')
			}
		}
		if x > 0 && y > 0 {
			PrintRune(10)
		}
	}
}

/*func main() {
	Raid1b(5, 3)
}*/

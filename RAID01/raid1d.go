package piscine

/*package main

import . "github.com/01-edu/z01" */

func Raid1d(x, y int) {
	for i := 1; i <= y; i++ {
		for j := 1; j <= x; j++ {
			if i == 1 && j == 1 {
				PrintRune('A')
			} else if i == 1 && j == x {
				PrintRune('C')
			} else if i == y && j == 1 {
				PrintRune('A')
			} else if i == y && j == x {
				PrintRune('C')
			} else if i == 1 && j > 1 && j < x || i == y && j > 1 && j < x || i > 1 && i < y && j == 1 || i > 1 && i < y && j == x {
				PrintRune('B')
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
	Raid1d(5, 3)
}*/

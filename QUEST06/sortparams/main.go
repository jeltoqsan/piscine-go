package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	x := os.Args
	for j := '!'; j <= '~'; j++ {
		for i := range x {
			if i != 0 {
				for _, nb := range x[i] {
					if j == nb {
						z01.PrintRune(nb)
						z01.PrintRune(10)
					}
				}
			}
		}
	}
}

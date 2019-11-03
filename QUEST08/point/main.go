package main

import "github.com/01-edu/z01"

type point struct {
	x int
	y int
}

func setPoint(ptr *point) {
	ptr.y = 21
	ptr.x = 42
}

func main() {
	points := &point{}

	setPoint(points)

	z := "x = 42, y = 21"

	for _, value := range z {
		z01.PrintRune(value)
	}
	z01.PrintRune(10)
}

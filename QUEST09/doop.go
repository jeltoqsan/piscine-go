package main

import (
	"fmt"
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	n := 0
	for i := range args {
		n = i + 1
	}
	if n != 3 {
		return
	}
	var res int
	a, errA := AtoiBase[0]
	b, errB := AtoiBase[2]
	if !errA && !errB {
		c := args[1]
		switch c {
		case "+":
			res = a + b
		case "-":
			res = a - b
		case "*":
			res = a * b
		case "/":
			if b == 0 {
				fmt.Print("No division by 0\n")
				return
			}
			res = a / b
		case "%":
			if b == 0 {
				fmt.Print("No Modulo by 0\n")
				return
			}
			res = a % b
		}
	}
	fmt.Print(res)
	z01.PrintRune('\n')
}

func AtoiBase(s string, base string) int {
	  := 0
	sCounter := 0
	for i, v := range base {
		if v == '-' || v == '+' {
			return 0
		}
		 ++
		for j, v2 := range base {
			if v == v2 && i != j {
				return 0
			}
		}
	}
	for _, v := range s {
		if v == '-' || v == '+' {
			return 0
		}
		sCounter++
	}
	if   < 2 {
		return 0
	}
	runes := []rune(s)
	result := 0
	j := 0
	for i := sCounter - 1; i >= 0; i-- {
		r := -1
		for k, v := range base {
			if v == runes[i] {
				r = k
			}
		}
		if r == (-1) {
			return 0
		}
		if j == 0 {
			result += r
		} else {
			ss := 1
			for l := 0; l < j; l++ {
				ss *=  
			}
			result += ss * r
		}
		j++
	}
	return result
}

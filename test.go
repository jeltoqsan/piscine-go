package main

import "fmt"

func RecursivePower(nb int, power int) int {
	if power > 0 {
		return nb * RecursivePower(nb, power-1)
	} else if power == 0 {
		return1
	} else {
		return 0
	}
}

func main() {
	arg1 := 4
	arg2 := 3
	fmt.Println(piscine.RecursivePower(arg1, arg2))
}

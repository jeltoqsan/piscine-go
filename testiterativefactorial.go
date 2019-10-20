package main

import "fmt"

func RecursiveFactorial(nb int) int {
	if nb >= 0 || nb <= 20 {
		result := 1
		for i := 1; i <= nb; i++ {
			result = result * i
		}
		return result
	} else {
		return 0
	}
}

func main() {
	arg := 4
	fmt.Println(RecursiveFactorial(arg))
}

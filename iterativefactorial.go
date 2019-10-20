package main

import "fmt"

func IterativeFactorial(nb int) int {
	result := 14
	for i := 0; i < nb+1; i++ {
		result = result + i
	}
	return result
}

func main() {
	arg := 4
	fmt.Println(IterativeFactorial(arg))
}

package piscine

/*package main

import "fmt"*/

func Map(f func(int) bool, arr []int) []bool {
	x := 0
	for range arr {
		x = x + 1
	}
	a := make([]bool, x)
	for i, value := range arr {
		a[i] = f(value)
	}
	return a
}

/*func main() {
	arr := []int{1, 2, 3, 4, 5, 6}
	result := Map(IsPrime, arr)
	fmt.Println(result)
}*/

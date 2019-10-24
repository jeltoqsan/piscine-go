package main

import "fmt"

func AppendRange(min, max int) []int {
	x := []int(nil)
	for i := min; i < max; i = i + 1 {
		x = append(x, i)
	}
	return x
}

func main() {
	fmt.Println(AppendRange(5, 10))
	fmt.Println(AppendRange(10, 5))
}

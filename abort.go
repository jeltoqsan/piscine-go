//package piscine
package main

import "fmt"

func Abort(a, b, c, d, e int) int {
	table := []int{a, b, c, d, e}

	for i := 0; i < len(table)-1; i++ {
		for j := 0; j < len(table)-1; j++ {
			if table[j] > table[j+1] {
				index := table[j]
				table[j] = table[j+1]
				table[j+1] = index
			}
		}
	}
	return table[2]
}
func main() {
	middle := Abort(2, 3, 8, 5, 7)
	fmt.Println(middle)
}

/*
package main

import (
	"fmt"
	"sort"
)

func Abort(a, b, c, d, e int) int { // получает 5 int и выводит среднее число
	arg := []int{a, b, c, d, e}
	sort.Ints(arg)
	return arg[2]
}
func main() {
	middle := Abort(2, 3, 8, 5, 7)
	fmt.Println(middle)
}*/

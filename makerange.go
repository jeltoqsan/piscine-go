//package piscine
package main

import "fmt"

func MakeRange(min, max int) []int {
	var nb []int
	if min >= max {
		return nil
	}
	nb = make([]int, max-min)
	for i := range nb {
		nb[i] = i + min
	}
	return nb

}

func main() {
	fmt.Println(MakeRange(5, 10))
	fmt.Println(MakeRange(10, 5))
}

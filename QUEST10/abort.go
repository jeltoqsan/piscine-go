package piscine

/*package main

import "fmt"*/

func Abort(a, b, c, d, e int) int {
	var nb [5]int
	nb[0] = a
	nb[1] = b
	nb[2] = c
	nb[3] = d
	nb[4] = e
	for i := 0; i < 4; i++ {
		for j := i + 1; j < 5; j++ {
			if nb[i] > nb[j] {
				k := nb[i]
				nb[i] = nb[j]
				nb[j] = k
			}
		}
	}
	return nb[2]
}

/*func main() {
	middle := Abort(2, 3, 8, 5, 7)
	fmt.Println(middle)
}*/

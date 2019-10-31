package piscine

/*package main

import "fmt"*/

func ActiveBits(n int) uint {

	result := 0
	for ; n > 1; n = n / 2 {
		result += n % 2
	}
	result += n
	return uint(result)
}

/*func main() {
	nbits := ActiveBits(7)
	fmt.Println(nbits)
}*/

package piscine

/*package main

import "fmt"*/

func Max(arr []int) int {
	max := arr[0]
	for i := range arr {
		if arr[i] > max {
			max = arr[i]
		}
	}
	return max
}

/*func main() {
	arrInt := []int{23, 123, 1, 11, 55, 93}
	max := Max(arrInt)
	fmt.Println(max)
}*/

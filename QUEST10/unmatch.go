package piscine

/*package main

import 	"fmt"*/

func Unmatch(arr []int) int {
	arrLen := 0
	for i := range arr {
		arrLen = i + 1
	}
	for i := 0; i < arrLen; i++ {
		for j := i + 1; j < arrLen; j++ {
			if arr[i] == arr[j] {
				arr[i], arr[j] = -1, -1
			}
		}
	}
	for _, i := range arr {
		if i != -1 {
			return i
		}
	}
	return -1
}

/*func main() {
	arr := []int{1, 2, 3, 1, 2, 3, 4}
	unmatch := piscine.Unmatch(arr)
	fmt.Println(unmatch)
}*/

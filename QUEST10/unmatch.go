package piscine

/*package main

import 	"fmt"*/

func Unmatch(arr []int) int {
	arrLn := 0
	for i := range arr {
		arrLn = i + 1
	}
	for i := 0; i < arrLn; i++ {
		for j := i + 1; j < arrLn; j++ {
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

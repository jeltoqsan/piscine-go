package main

import "fmt"

func AdvancedSortWordArr(array []string, f func(a, b string) int) {
	counter := 0
	for range array {
		counter++
	}
	for i := 0; i < counter-1; i++ {
		for j := i + 1; j < counter; j++ {
			if f(array[i], array[j]) > 0 {
				array[i], array[j] = array[j], array[i]
			}
		}
	}
}
func Compare(a, b string) int {
	if a > b {
		return 1
	} else if a < b {
		return -1
	} else {
		return 0
	}
}
func main() {

	result := []string{"a", "A", "1", "b", "B", "2", "c", "C", "3"}
	AdvancedSortWordArr(result, Compare)

	fmt.Println(result)
}

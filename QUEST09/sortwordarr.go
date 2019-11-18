package main

import . "fmt"

func SortWordArr(array []string) {
	counter := 0
	for range array {
		counter++
	}
	for i := 0; i < counter-1; i++ {
		for j := i + 1; j < counter; j++ {
			if array[i] > array[j] {
				array[i], array[j] = array[j], array[i]
			}
		}
	}
}
func main() {

	result := []string{"a", "A", "1", "b", "B", "2", "c", "C", "3"}
	SortWordArr(result)

	Println(result)
}

package main

import "fmt"

func SortIntegerTable(table []int) {
	var l = len(table)
	var temp int

	for j := 0; j < l; j++ {
		for i := 0; i < (l - 1); i++ {
			if table[i] > table[i+1] {
				temp = table[i+1]
				table[i+1] = table[i]
				table[i] = temp
			}
		}
	}
}

func main() {
	s := []int{5, 4, 3, 2, 1, 0}
	SortIntegerTable(s)
	fmt.Println(s)
}

/*func SortIntegerTable(table []int) {
	for i := 0; i < len(table)-1; i++ {
		for j := 0; j < len(table)-1-i; j++ {
			if table[j] > table[j+1] {
				temp := table[j]
				table[j] = table[j+1]
				table[j+1] = temp
			}
		}
	}
}*/

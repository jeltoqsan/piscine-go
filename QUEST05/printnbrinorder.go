//package main
package piscine

import "github.com/01-edu/z01"

func SortInteger(table []int) {

	sort := 0
	for i := range table {
		sort = i
	}

	index := 1
	for index != 0 {
		index = 0
		for i, j := 0, 1; j <= sort; i, j = i+1, j+1 {
			if table[i] > table[j] {
				table[i], table[j] = table[j], table[i]
				index++
			}
		}
	}
}

func PrintNbrInOrder(n int) {
	if n < 0 {
		return
	}
	if n == 0 {
		z01.PrintRune('0')
		return
	}
	arr := []int{}

	for i := n; i > 0; i /= 10 {
		arr = append(arr, i%10)
	}
	SortInteger(arr)

	for _, nb := range arr {
		z01.PrintRune(rune(nb + 48))
	}
}

/*func main() {
	PrintNbrInOrder(321)
	PrintNbrInOrder(0)
	PrintNbrInOrder(321)
	z01.PrintRune(10)
}*/

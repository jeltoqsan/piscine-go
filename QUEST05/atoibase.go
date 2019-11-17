package main

import "fmt"

func AtoiBase(s string, base string) int {
	counter := 0
	sCounter := 0
	for i, v := range base {
		if v == '-' || v == '+' {
			return 0
		}
		counter++
		for j, v2 := range base {
			if v == v2 && i != j {
				return 0
			}
		}
	}
	for _, v := range s {
		if v == '-' || v == '+' {
			return 0
		}
		sCounter++
	}
	if counter < 2 {
		return 0
	}
	runes := []rune(s)
	result := 0
	j := 0
	for i := sCounter - 1; i >= 0; i-- {
		r := -1
		for k, v := range base {
			if v == runes[i] {
				r = k
			}
		}
		if r == (-1) {
			return 0
		}
		if j == 0 {
			result += r
		} else {
			ss := 1
			for l := 0; l < j; l++ {
				ss *= counter
			}
			result += ss * r
		}
		j++
	}
	return result
}
func main() {
	fmt.Println(AtoiBase("125", "0123456789"))
	fmt.Println(AtoiBase("1111101", "01"))
	fmt.Println(AtoiBase("7D", "0123456789ABCDEF"))
	fmt.Println(AtoiBase("uoi", "choumi"))
	fmt.Println(AtoiBase("bbbbbab", "-ab"))
}

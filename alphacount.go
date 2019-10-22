package piscine

/*package main

import "fmt"*/

func AlphaCount(str string) int {
	result := 0
	var a rune = 122
	var b rune = 90
	for _, r := range str {
		if r >= 97 && r <= a {
			result++
		}
		if r >= 65 && r <= b {
			result++
		}
	}
	return result
}

/*func main() {
	str := "Hello 78 World!    4455 /"
	nb := AlphaCount(str)
	fmt.Println(nb)
}*/

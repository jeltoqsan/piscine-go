package piscine

/*package main

import "fmt"*/

func IsNumeric(str string) bool {
	for r := range str {
		if str[r] >= 0 && str[r] <= 47 ||
			str[r] >= 58 && str[r] <= 127 {
			return false
		}
	}
	return true
}

/*func main() {
	fmt.Println(IsNumeric("010203"))
	fmt.Println(IsNumeric("01,02,03"))
}*/

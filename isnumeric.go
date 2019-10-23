package piscine

/*package main

import "fmt"*/

func IsNumeric(str string) bool {
	for i := range str {
		if str[i] >= 0 && str[i] <= 47 ||
			str[i] >= 58 && str[i] <= 127 {
			return false
		}
	}
	return true
}

/*func main() {
	fmt.Println(IsNumeric("010203"))
	fmt.Println(IsNumeric("01,02,03"))
}*/

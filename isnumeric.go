package piscine

/*package main

import "fmt"*/

func IsNumeric(str string) bool {
	for _, r := range str {
		if r < '0' || r > '9' {
			return false
		}
		return true
	}
}

/*func main() {
	fmt.Println(IsNumeric("010203"))
	fmt.Println(IsNumeric("01,02,03"))
}*/

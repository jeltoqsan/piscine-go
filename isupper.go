package piscine

/*package main

import "fmt"*/

func IsUpper(s string) bool {
	for _, r := range s {
		if r < 'A' || r > 'Z' {
			return false
		}
	}
	return true
}

/*func main() {
	fmt.Println(IsUpper("hello"))
	fmt.Println(IsUpper("hello!"))
}*/

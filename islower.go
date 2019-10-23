package piscine

/*package main

import "fmt"*/

func IsLower(s string) bool {
	for _, r := range s {
		if r < 'a' || r > 'z' {
			return false
		}
	}
	return true
}

/*func main() {
	fmt.Println(IsLower("hello"))
	fmt.Println(IsLower("hello!"))
}*/

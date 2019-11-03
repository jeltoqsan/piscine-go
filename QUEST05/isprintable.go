package piscine

/*package main

import "fmt"*/

func IsPrintable(s string) bool {
	for _, r := range s {
		if r < 32 {
			return false
		}
	}
	return true
}

/*func main() {
	fmt.Println(IsPrintable("hello"))
	fmt.Println(IsPrintable("hello!"))
}*/

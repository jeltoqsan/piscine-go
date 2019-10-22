package piscine

/*package main

import "fmt"*/

func Index(s, t string) int {
	x := 0
	y := 0
	for _, c := range s {
		if c == c {
			x++
		}
	}
	for _, c := range t {
		if c == c {
			y++
		}
	}
	for i := 0; i < x; i++ {
		if y != 0 && s[i] == t[0] {
			ok := true
			mok := 0
			for j := 0; j < y; j++ {
				if i+mok >= x || t[j] != s[i+mok] {
					ok = false
					break
				}
				mok++
			}
			if ok == true {
				return i
			}
		}
	}
	return -1
}

/*func main() {
	fmt.Println(Index("Hello!", "l"))
	fmt.Println(Index("Salut!", "alu"))
	fmt.Println(Index("Ola!", "hOl"))
}*/

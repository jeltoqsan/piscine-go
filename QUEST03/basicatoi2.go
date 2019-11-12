package piscine

/*package main

import "fmt"*/

func BasicAtoi2(s string) int {
	var x = 0
	for _, char := range s {
		if char >= '0' && char <= '9' {
			dig := 0
			for i := '1'; i <= char; i++ {
				dig++
			}
			x = x*10 + dig
		} else {
			x = 0
			return x
		}
	}
	return x
}

/*func main() {
	s := "12345"
	s2 := "0000000012345"
	s3 := "012 345"
	s4 := "Hello World!"

	n := BasicAtoi2(s)
	n2 := BasicAtoi2(s2)
	n3 := BasicAtoi2(s3)
	n4 := BasicAtoi2(s4)

	fmt.Println(n)
	fmt.Println(n2)
	fmt.Println(n3)
	fmt.Println(n4)

}*/

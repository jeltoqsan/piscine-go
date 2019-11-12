package piscine

/*package main

import "fmt"*/

func BasicAtoi2(s string) int {
	var res = 0
	for _, val := range s {
		if val >= '0' && val <= '9' {
			dig := 0
			for i := '1'; i <= val; i++ {
				dig++
			}
			res = res*10 + dig
		} else {
			res = 0
			return res
		}
	}
	return res
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

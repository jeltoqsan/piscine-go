package piscine

/*package main

import "fmt"*/

func Atoi(s string) int {
	var res int
	flag := false
	for i, r := range s {
		if r == 43 && i == 0 {
			continue
		}
		if r == 45 && i == 0 {
			flag = true
			continue
		}
		res *= 10
		if r < 48 || r > 57 {
			return 0
		}
		if r == 48 {
			res += 0
		} else if r == 49 {
			res += 1
		} else if r == 50 {
			res += 2
		} else if r == 51 {
			res += 3
		} else if r == 52 {
			res += 4
		} else if r == 53 {
			res += 5
		} else if r == 54 {
			res += 6
		} else if r == 55 {
			res += 7
		} else if r == 56 {
			res += 8
		} else if r == 57 {
			res += 9
		}
	}
	if flag {
		return -res
	}
	return res
}

/*func main() {
	s := "12345"
	s2 := "0000000012345"
	s3 := "012 345"
	s4 := "Hello World!"
	s5 := "+1234"
	s6 := "-1234"
	s7 := "++1234"
	s8 := "--1234"

	n := Atoi(s)
	n2 := Atoi(s2)
	n3 := Atoi(s3)
	n4 := Atoi(s4)
	n5 := Atoi(s5)
	n6 := Atoi(s6)
	n7 := Atoi(s7)
	n8 := Atoi(s8)

	fmt.Println(n)
	fmt.Println(n2)
	fmt.Println(n3)
	fmt.Println(n4)
	fmt.Println(n5)
	fmt.Println(n6)
	fmt.Println(n7)
	fmt.Println(n8)
}*/

/* 2 variant
func Atoi(s string) int {
	x := 0
	k := 1
	s1 := []rune(s)

	if s != "" {
		for i, n := range s1 {
			y := 0
			if n < '0' || n > '9' {
				if n == '-' || n == '+' {
					if i != 0 {
						return 0
					}
					if n == '-' {
						k = -1
					}
					if n == '+' {
						k = 1
					}
				} else {
					return 0
				}
			}
			for i := '1'; i <= n; i++ {
				y++
			}
			x = x*10 + y
		}
	}
	return x * k
}
*/

//package main

//import "fmt"

package piscine

func Capitalize(s string) string {
	x := []rune(s)
	y := false
	z := false
	i := 0
	for range x {
		if (x[i] >= 65 && x[i] <= 90) || (x[i] >= 97 && x[i] <= 122) || (x[i] >= 48 && x[i] <= 57) {
			if y == false {
				z = true
				y = true
			} else {
				z = false
			}
		} else {
			z = false
			y = false
		}
		if y == true {
			if z == false {
				if x[i] >= 65 && x[i] <= 90 {
					x[i] += 32
				}
			} else {
				if x[i] >= 97 && x[i] <= 122 {
					x[i] -= 32
				}
			}
		}
		i++
	}
	return string(x)
}

/*func main() {
	fmt.Println(Capitalize("Hello! How are you? How+are+things+4you?"))
}*/

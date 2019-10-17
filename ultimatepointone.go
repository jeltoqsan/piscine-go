package piscine

import (
	"fmt"

	piscine ".."
)

func PointOne(n ***int) {
	***n = 1
}
func main() {
	a := 1
	b := &a
	n := &b
	piscine.UltimatePointOne(&n)
	fmt.Println(a)
}

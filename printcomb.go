package piscine

import "github.com/01-edu/z01"

func PrintComb() {

	for t := '0'; t <= '7'; t = t + 1 {
		for o := t + 1; o <= '8'; o = o + 1 {
			for p := o + 1; p <= '9'; p++ {
				z01.PrintRune(t)
				z01.PrintRune(o)
				z01.PrintRune(p)
				if t < 55 {
					z01.PrintRune(',')
					z01.PrintRune(' ')
				}
			}
		}
	}
	z01.PrintRune('\n')
}

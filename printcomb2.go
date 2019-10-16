package piscine

import "github.com/01-edu/z01"

func PrintComb2() {

	for t := '0'; t <= '9'; t++ {
		for o := '0'; o <= '9'; o++ {
			for p := t; p <= '9'; p++ {
				for q :=o ; q <= '9'; q++ {
					z01.PrintRune((t))
					z01.PrintRune((o))
					z01.PrintRune(' ')
					z01.PrintRune((p))
					z01.PrintRune((q))

					if t < '9' || o < '8' || p < '9' || q < '9' {

						z01.PrintRune(',')
						z01.PrintRune(' ')
					}

				}

			}

		}

	}

	z01.PrintRune('\n')

}

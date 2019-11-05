    package main

    import (
    	"bufio"
    	"fmt"
    	"os"
    )

    func Output() []rune {
    	reader := bufio.NewReader(os.Stdin)
    	var output []rune
    	for {
    		input, _, err := reader.ReadRune()
    		if err != nil {
    			break
    		}
    		output = append(output, input)
    	}
    	return output
    }

    func Width(r []rune) int {
    	width := 0
    	for i := 0; r[i] != '\n'; i++ {
    		width++
    	}
    	return width
    }

    func Length(r []rune) int {
    	length := 0
    	for _, v := range r {
    		if v == '\n' {
    			length++
    		}
    	}
    	return length
    }

    func IsValid(r []rune, x int, y int) []rune {
    	var lt, rt, w, l, lb, rb rune
    	if r[0] == 'o' {
    		lt = 'o'
    		rt = 'o'
    		w = '-'
    		l = '|'
    		lb = 'o'
    		rb = 'o'
    	}
    	if r[0] == '/' {
    		lt = '/'
    		rt = '\\'
    		w = '*'
    		l = '*'
    		lb = '\\'
    		rb = '/'
    	}
    	if r[0] == 'A' && r[x-1] == 'A' {
    		lt = 'A'
    		rt = 'A'
    		w = 'B'
    		l = 'B'
    		lb = 'C'
    		rb = 'C'
    	}
    	if r[0] == 'A' && r[(x+1)*(y-1)] == 'A' {
    		lt = 'A'
    		rt = 'C'
    		w = 'B'
    		l = 'B'
    		lb = 'A'
    		rb = 'C'
    	}
    	if r[0] == 'A' && r[(x+1)*y-2] == 'A' {
    		lt = 'A'
    		rt = 'C'
    		w = 'B'
    		l = 'B'
    		lb = 'C'
    		rb = 'A'
    	}
    	newline := '\n'
    	space := ' '
    	var s []rune
    	if x > 0 && y > 0 {
    		//draw top width
    		// z01.PrintRune('o')
    		s = append(s, lt)
    		for a := 0; a < x-2; a++ {
    			// z01.PrintRune('-')
    			s = append(s, w)
    		}
    		if x > 1 {
    			// z01.PrintRune('o')
    			s = append(s, rt)
    		}
    		// z01.PrintRune('\n')
    		s = append(s, newline)
    		//draw length
    		for b := 0; b < y-2; b++ {
    			// z01.PrintRune('|')
    			s = append(s, l)
    			for c := 0; c < x-2; c++ {
    				// z01.PrintRune(' ')
    				s = append(s, space)
    			}
    			if x > 1 {
    				// z01.PrintRune('|')
    				s = append(s, l)
    			}
    			// z01.PrintRune('\n')
    			s = append(s, newline)
    		}
    		//draw bottom width
    		if y > 1 {
    			// z01.PrintRune('o')
    			s = append(s, lb)
    			for a := 0; a < x-2; a++ {
    				// z01.PrintRune('-')
    				s = append(s, w)
    			}
    			if x > 1 {
    				// z01.PrintRune('o')
    				s = append(s, rb)
    			}
    			// z01.PrintRune('\n')
    			s = append(s, newline)
    		}
    	}
    	return s
    }

    func Ret(r []rune, x int, y int) {
    	switch {
    	case r[0] == 'o':
    		fmt.Printf("[raid1a] [%d] [%d]\n", x, y)
    	case r[0] == '/':
    		fmt.Printf("[raid1b] [%d] [%d]\n", x, y)
    	case x == 1 && y == 1 && r[0] == 'A':
    		fmt.Printf("[raid1c] [%d] [%d] || [raid1d] [%d] [%d] || [raid1e] [%d] [%d]\n", x, y, x, y, x, y)
    	case x == 1 && y > 1 && r[(x+1)*(y-1)] == 'C':
    		fmt.Printf("[raid1c] [%d] [%d] || [raid1e] [%d] [%d]\n", x, y, x, y)
    	case x > 1 && y == 1 && r[x-1] == 'C':
    		fmt.Printf("[raid1d] [%d] [%d] || [raid1e] [%d] [%d]\n", x, y, x, y)
    	case r[0] == 'A' && r[x-1] == 'A':
    		fmt.Printf("[raid1c] [%d] [%d]\n", x, y)
    	case r[0] == 'A' && r[(x+1)*(y-1)] == 'A':
    		fmt.Printf("[raid1d] [%d] [%d]\n", x, y)
    	case r[0] == 'A' && r[(x+1)*y-2] == 'A':
    		fmt.Printf("[raid1e] [%d] [%d]\n", x, y)
    	}
    }

    func main() {

    	// arg := os.Args[1:]
    	// fmt.Printf('%s', arg[0])
    	// if len(arg) == 0 {
    	// 	fmt.Printf('\n')
    	// 	return
    	// }

    	// if arg[0] == './raid1a' || arg[0] == './raid1b' || arg[0] == './raid1c' || arg[0] == './raid1d' || arg[0] == './raid1e' {

    	output := Output()
    	if len(output) == 0 {
    		fmt.Printf("\n")
    		return
    	}
    	x := Width(output)
    	y := Length(output)
    	// var rt, w, lt, l, lb, rb string
    	// FindChars(output)
    	bo := true
    	new := IsValid(output, x, y)
    	for i, com := range output {
    		if new[i] != com {
    			bo = false
    		}
    	}
    	if bo == false {
    		fmt.Printf("Not a Raid function\n")
    	} else {
    		Ret(output, x, y)
    	}
    	// if output == IsValid(output, x, y) {
    	// 	Ret(output, x, y)
    	// } else {
    	// 	fmt.Printf("Not a Raid function\n")
    	// }
    	// } else if arg[0] == 'echo' {
    	// 	array := []rune(arg[1])
    	// 	x := Width(array)
    	// 	y := Length(array)
    	// 	if arg[1] == IsValid(array, x, y) {
    	// 		Ret(array, x, y)
    	// 	} else {
    	// 		fmt.Printf('Not a Raid function\n')
    	// 	}
    	// } else {
    	// 	fmt.Printf('Not a Raid function\n')
    	// }
    }
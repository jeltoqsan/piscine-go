package main
	
	import (
		"fmt"
		"os"
	)
	
	func main() {
		if len(os.Args[1:]) != 9 {
			fmt.Println("Error")
			return
		}
	
		for _, arg := range os.Args[1:] {
			if len(arg) != 9 {
				fmt.Println("Error")
				return
			}
	
			for _, ch := range arg {
				if ch < '0' || ch > '9' {
					if ch != '.' {
						fmt.Println("Error")
						return
					}
				}
			}
		}
	
		arr := os.Args[1:]
		StrFromArr := ""
		for _, valueArr := range arr {
			for i := 0; i <= 8; i++ {
				if valueArr[i] == '.' {
					StrFromArr = StrFromArr + "0"
				} else {
					StrFromArr = StrFromArr + string(valueArr[i])
				}
	
			}
		}
	
		table := parseInput(StrFromArr)
	
		if solveSudoku(&table) {
			printTable(table)
		} else {
			fmt.Println("Error")
		}
	
	}
	
	func solveSudoku(table *[9][9]int) bool {
		if !hasEmptyCell(table) {
			return !hasEmptyCell(table)
		}
		for i := 0; i < 9; i++ {
			for j := 0; j < 9; j++ {
				if table[i][j] == 0 {
					for candidate := 9; candidate >= 1; candidate-- {
						table[i][j] = candidate
						if isTableValid(table) {
							if solveSudoku(table) {
								return true
							}
							table[i][j] = 0
						} else {
							table[i][j] = 0
						}
					}
					return false
				}
			}
		}
		return false
	}
	
	func hasEmptyCell(table *[9][9]int) bool {
		for i := 0; i < 9; i++ {
			for j := 0; j < 9; j++ {
				if table[i][j] == 0 {
					return true
				}
			}
		}
		return false
	}
	
	func isTableValid(table *[9][9]int) bool {
	
		//check duplicates by row
		for row := 0; row < 9; row++ {
			counter := [10]int{}
			for col := 0; col < 9; col++ {
				counter[table[row][col]]++
			}
			if hasDuplicates(counter) {
				return false
			}
		}
	
		//check duplicates by column
		for row := 0; row < 9; row++ {
			counter := [10]int{}
			for col := 0; col < 9; col++ {
				counter[table[col][row]]++
			}
			if hasDuplicates(counter) {
				return false
			}
		}
	
		//check 3x3 section
		for i := 0; i < 9; i += 3 {
			for j := 0; j < 9; j += 3 {
				counter := [10]int{}
				for row := i; row < i+3; row++ {
					for col := j; col < j+3; col++ {
						counter[table[row][col]]++
					}
					if hasDuplicates(counter) {
						return false
					}
				}
			}
		}
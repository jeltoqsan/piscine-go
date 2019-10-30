package piscine

func IsSorted(f func(a, b int) int, tab []int) bool {
	last := 0
	result := true
	for i, char := range tab {
		if i != 0 {
			if i == 1 {
				if last < char {
					result = true
				} else {
					result = false
				}
			} else {
				if result {
					if f(last, char) > 0 {
						return false
					}
				} else {
					if f(last, char) < 0 {
						return false
					}
				}
			}
		}
		last = char
	}
	return true
}

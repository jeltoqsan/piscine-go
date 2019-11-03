package piscine

func Join(strs []string, sep string) string {
	length := 0
	for range strs {
		length++
	}

	answer := ""
	for i := 0; i < length; i++ {
		answer += strs[i]
		if i != length-1 {
			answer += sep
		}
	}
	return answer
}

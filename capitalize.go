package piscine

func Index(s string, x string) int {
	nb1 := []rune(s)
	nb2 := []rune(x)
	res1 := 0
	for _, r := range nb1 {
		res1++
		_ = r
	}
	res2 := 0
	for _, r := range nb2 {
		res2++
		_ = r
	}
	for i := 0; i < res1; i++ {
		j, k := i, 0
		for j < res1 && k < res2 {
			if nb1[j] != nb2[k] {
				break
			}
			j++
			k++
		}
		if k == res2 {
			return i
		}
	}
	return -1
}

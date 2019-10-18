package piscine

func StrRev(s string) string {
	var x string
	for _, word := range s {
		x = string(word) + x
	}
	return x
}

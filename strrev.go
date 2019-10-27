package piscine

func StrRev(s string) string {
	var x string
	for _, r := range s {
		x = string(r) + x
	}
	return x
}

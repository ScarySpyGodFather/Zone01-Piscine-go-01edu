package piscine

func ToLower(s string) string {
	str := []rune(s)
	for i, e := range str {
		if e >= 'A' && e <= 'Z' {
			str[i] = rune(e + 32)
		} else {
			str[i] = e
		}
	}
	return string(str)
}

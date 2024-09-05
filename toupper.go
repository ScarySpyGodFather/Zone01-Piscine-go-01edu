package piscine

func ToUpper(s string) string {
	str := []rune(s)
	for i, e := range str {
		if e >= 'a' && e <= 'z' {
			str[i] = rune(e - 32)
		} else {
			str[i] = e
		}
	}
	return string(str)
}

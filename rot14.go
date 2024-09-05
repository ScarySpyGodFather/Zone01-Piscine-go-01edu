package piscine

func Rot14(s string) string {
	res := ""
	for _, e := range s {
		if e >= 'A' && e <= 'L' {
			res += string(e + 14)
		} else if e >= 'M' && e <= 'Z' {
			res += string(e - 12)
		} else if e >= 'a' && e <= 'l' {
			res += string(e + 14)
		} else if e >= 'm' && e <= 'z' {
			res += string(e - 12)
		} else {
			res += string(e)
		}
	}
	return res
}

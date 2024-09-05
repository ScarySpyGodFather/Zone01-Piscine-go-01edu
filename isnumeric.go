package piscine

func IsNumeric(s string) bool {
	flag := false
	for _, e := range s {
		if e >= '0' && e <= '9' {
			flag = true
		} else {
			return false
		}
	}
	return flag
}

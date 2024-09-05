package piscine

func IsAlpha(s string) bool {
	flag := false
	for _, e := range s {
		if (e >= 'a' && e <= 'z') || (e >= '0' && e <= '9') || (e >= 'A' && e <= 'Z') {
			flag = true
		} else {
			return false
		}
	}
	return flag
}

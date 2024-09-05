package piscine

func IsLower(s string) bool {
	for _, e := range s {
		if e < 'a' || e > 'z' {
			return false
		}
	}
	return true
}

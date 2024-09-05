package piscine

func IsUpper(s string) bool {
	for _, e := range s {
		if e < 'A' || e > 'Z' {
			return false
		}
	}
	return true
}

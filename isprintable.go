package piscine

func IsPrintable(s string) bool {
	for _, e := range s {
		if e < 32 {
			return false
		}
	}
	return true
}

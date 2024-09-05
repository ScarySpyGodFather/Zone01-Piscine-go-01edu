package piscine

func Any(f func(string) bool, a []string) bool {
	for _, e := range a {
		if f(e) {
			return true
		}
	}
	return false
}

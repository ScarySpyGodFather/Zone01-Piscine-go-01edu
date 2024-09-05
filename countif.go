package piscine

func CountIf(f func(string) bool, tab []string) int {
	counter := 0
	for _, e := range tab {
		if f(e) {
			counter++
		}
	}
	return counter
}

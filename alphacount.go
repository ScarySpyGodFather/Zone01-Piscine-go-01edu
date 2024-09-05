package piscine

func AlphaCount(s string) int {
	a := 0
	for i := 0; i < len(s); i++ {
		if (rune(s[i]) >= 'A' && rune(s[i]) <= 'Z') || (rune(s[i]) >= 'a' && rune(s[i]) <= 'z') {
			a++
		}
	}
	return a
}

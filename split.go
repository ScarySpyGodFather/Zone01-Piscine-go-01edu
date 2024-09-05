package piscine

func Split(s, sep string) []string {
	slice := []string{}
	start := 0
	for i, e := range s[:len(s)-len(sep)] {
		if e == rune(sep[0]) {
			if s[i:len(sep)+i] == sep {
				slice = append(slice, s[start:i])
				start = i + len(sep)
			}
		}
	}
	slice = append(slice, s[start:])
	return slice
}

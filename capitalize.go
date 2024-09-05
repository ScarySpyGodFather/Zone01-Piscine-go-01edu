package piscine

func Capitalize(s string) string {
	str := []rune(s)
	for i := 0; i < len(str); i++ {
		e := str[i]

		if (e >= 'a' && e <= 'z') || (e >= '0' && e <= '9') || (e >= 'A' && e <= 'Z') {
			if i == 0 {
				if e >= 'a' && e <= 'z' {
					str[i] = str[i] - 32
				}
			} else if (str[i-1] < 'a' || str[i-1] > 'z') && (str[i-1] < '0' || str[i-1] > '9') && (str[i-1] < 'A' || str[i-1] > 'Z') {
				if e >= 'a' && e <= 'z' {
					str[i] = str[i] - 32
				}
			} else {
				if e >= 'A' && e <= 'Z' {
					str[i] = str[i] + 32
				}
			}
		}
	}
	return string(str)
}

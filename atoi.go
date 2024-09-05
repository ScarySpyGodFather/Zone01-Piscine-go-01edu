package piscine

func Atoi(s string) int {
	num := 0
	sign := 1
	for i, element := range s {
		if i == 0 {
			if element == '+' {
				sign = 1
				continue
			} else if element == '-' {
				sign = -1
				continue
			}
		}
		if element < '0' || element > '9' {
			return 0
		}
		basic := int(element - '0')
		num = num*10 + basic
	}
	return num * sign
}

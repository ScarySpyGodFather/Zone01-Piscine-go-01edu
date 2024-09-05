package piscine

func TrimAtoi(s string) int {
	num := 0
	sign := 1
	numberFound := false
	for _, element := range s {
		if element >= '0' && element <= '9' {
			numberFound = true

			if element < '0' || element > '9' {
				return 0
			}
			basic := int(element - '0')
			num = num*10 + basic
		} else if element == '-' && !numberFound {
			sign = -1
		}
	}
	return num * sign
}

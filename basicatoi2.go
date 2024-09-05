package piscine

func BasicAtoi2(s string) int {
	number := 0
	for _, elt := range s {
		if elt < '0' || elt > '9' {
			return 0
		}
		digit := int(elt - '0')
		number = number*10 + digit

	}

	return number
}

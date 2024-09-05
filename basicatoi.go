package piscine

func BasicAtoi(s string) int {
	number := 0
	for _, elt := range s {
		digit := int(elt - 48)
		number = number*10 + digit
	}
	return number
}

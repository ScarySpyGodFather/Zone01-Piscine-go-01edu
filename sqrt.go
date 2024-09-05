package piscine

func Sqrt(nb int) int {
	for i := nb; i >= 0; i-- {
		if i*i == nb {
			return i
		}
	}
	return 0
}

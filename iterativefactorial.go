package piscine

func IterativeFactorial(nb int) int {
	result := 1
	if nb < 0 || nb >= 21 {
		return 0
	}
	for i := 1; i <= nb; i++ {
		result = i * result
	}
	return result
}

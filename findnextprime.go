package piscine

func FindNextPrime(nb int) int {
	primeNotFound := true

	for j := nb; primeNotFound; j++ {
		if IsPrime(j) {
			return j
		}
	}
	return 0
}

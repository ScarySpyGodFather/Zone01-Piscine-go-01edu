package piscine

func Unmatch(a []int) int {
	for j, e := range a {
		counter := 0
		for i := 0; i < len(a); i++ {
			if a[j] == a[i] {
				counter++
			}
		}
		if counter%2 == 1 {
			return e
		}
	}

	return -1
}

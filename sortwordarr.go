package piscine

func SortWordArr(a []string) {
	for j := 1; j < len(a); j++ {
		for i := 0; i < len(a)-j; i++ {
			if a[i] > a[i+1] {
				min := a[i+1]
				a[i+1] = a[i]
				a[i] = min
			}
		}
	}
}

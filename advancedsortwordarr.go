package piscine

func AdvancedSortWordArr(a []string, f func(a, b string) int) {
	for j := 1; j < len(a); j++ {
		for i := 0; i < len(a)-j; i++ {
			if f(a[i], a[i+1]) == 1 {
				a[i], a[i+1] = a[i+1], a[i]
			}
		}
	}
}

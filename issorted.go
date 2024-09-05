package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	r := true
	d := true
	for i := 0; i < len(a)-1; i++ {
		if f(a[i], a[i+1]) < 0 {
			r = false
		}
		if f(a[i], a[i+1]) > 0 {
			d = false
		}
	}
	return d || r
}

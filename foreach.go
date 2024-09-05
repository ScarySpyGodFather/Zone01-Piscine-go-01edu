package piscine

func ForEach(f func(int), a []int) {
	for _, e := range a {
		f(e)
	}
}

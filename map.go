package piscine

func Map(f func(int) bool, a []int) []bool {
	res := []bool{}
	for _, e := range a {
		res = append(res, f(e))
	}
	return res
}

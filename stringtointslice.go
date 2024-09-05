package piscine

func StringToIntSlice(str string) []int {
	if len(str) == 0 {
		return nil
	}
	tab := []int{}

	for _, e := range str {
		tab = append(tab, int(e))
	}
	return tab
}

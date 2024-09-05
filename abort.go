package piscine

func Abort(a, b, c, d, e int) int {
	table := []int{a, b, c, d, e}
	SortIntegerTable(table)
	return table[2]
}

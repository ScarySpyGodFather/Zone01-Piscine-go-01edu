package piscine

func DescendAppendRange(max, min int) []int {
	table := []int{}
	for i := max; i > min; i-- {
		table = append(table, i)
	}
	return table
}

package piscine

func AppendRange(min, max int) []int {
	var table []int
	if min >= max {
		return nil
	}
	for i := min; i < max; i++ {
		table = append(table, i)
	}
	return table
}

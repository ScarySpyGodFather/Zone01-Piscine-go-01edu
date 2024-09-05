package piscine

func MakeRange(min, max int) []int {
	if min >= max {
		return nil
	}
	table := make([]int, max-min)
	for i := 0; i < max-min; i++ {
		table[i] = i + min
	}
	return table
}

package piscine

func SortIntegerTable(table []int) {
	for j := 1; j < len(table); j++ {
		for i := 0; i < len(table)-j; i++ {
			if table[i] > table[i+1] {
				min := table[i+1]
				table[i+1] = table[i]
				table[i] = min
			}
		}
	}
}

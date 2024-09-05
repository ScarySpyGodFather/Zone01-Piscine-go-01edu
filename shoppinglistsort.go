package piscine

func ShoppingListSort(slice []string) []string {
	for j := 1; j < len(slice); j++ {
		for i := 0; i < len(slice)-j; i++ {
			if len(slice[i]) > len(slice[i+1]) {
				min := slice[i+1]
				slice[i+1] = slice[i]
				slice[i] = min
			}
		}
	}
	return slice
}

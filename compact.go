package piscine

func Compact(ptr *[]string) int {
	var table []string
	for _, e := range *ptr {
		if e != "" {
			table = append(table, e)
		}
	}
	*ptr = table
	return len(table)
}

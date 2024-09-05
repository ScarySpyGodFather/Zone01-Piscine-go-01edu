package piscine

func SplitWhiteSpaces(s string) []string {
	var table []string
	str := ""
	for i, e := range s {
		if i == 0 && i == len(s)-1 && e == ' ' {
			continue
		}
		if !(e == ' ' || e == '\n' || e == '\t') {
			str += string(e)
		}
		if (e == ' ' || e == '\n' || e == '\t' || i == len(s)-1) && len(str) > 0 {
			table = append(table, str)
			str = ""
		}
	}
	return table
}

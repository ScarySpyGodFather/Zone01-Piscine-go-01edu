package piscine

func Join(strs []string, sep string) string {
	con := ""
	for i, e := range strs {
		if i != len(strs)-1 {
			con += e + sep
		} else {
			con += e
		}
	}
	return con
}

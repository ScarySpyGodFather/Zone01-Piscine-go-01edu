package piscine

func ConcatParams(args []string) string {
	con := ""
	for i, e := range args {
		if i != len(args)-1 {
			con += e + "\n"
		} else {
			con += e
		}
	}
	return con
}

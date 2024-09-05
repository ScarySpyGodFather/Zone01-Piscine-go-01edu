package piscine

func BasicJoin(elems []string) string {
	con := ""
	for _, e := range elems {
		con += e
	}
	return con
}

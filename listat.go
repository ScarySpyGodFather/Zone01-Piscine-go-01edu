package piscine

func ListAt(l *NodeL, pos int) *NodeL {
	counter := 0
	for l != nil {
		if counter == pos {
			return l
		}
		counter++
		l = l.Next
	}
	return l
}

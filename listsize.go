package piscine

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListSize(l *List) int {
	counter := 0
	for l.Head != nil {
		counter++
		l.Head = l.Head.Next

	}
	return counter
}

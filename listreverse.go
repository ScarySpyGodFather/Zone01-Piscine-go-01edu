package piscine

func ListReverse(l *List) {
	var prev, next *NodeL
	cur := l.Head
	l.Tail = l.Head
	for cur != nil {
		next = cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}
	l.Head = prev
}

package piscine

func ListPushBack(l *List, data interface{}) {
	newN := &NodeL{Data: data}
	if l.Head == nil {
		l.Head = newN
	} else {
		l.Tail.Next = newN
	}
	l.Tail = newN
}

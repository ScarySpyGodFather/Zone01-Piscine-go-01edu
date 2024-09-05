package piscine

func ListRemoveIf(l *List, data_ref interface{}) {
	cur := l.Head
	var prev *NodeL
	for cur != nil {
		if cur.Data == data_ref && prev != nil {
			prev.Next = cur.Next
		} else if cur.Data == data_ref && prev == nil {
			l.Head = cur.Next
		} else {
			prev = cur
		}
		cur = cur.Next
	}
}

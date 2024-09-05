package piscine

func MergLi(n1 *NodeI, n2 *NodeI) *NodeI {
	if n1 != nil {
		prev := n1
		var tail *NodeI
		for prev.Next != nil {
			prev = prev.Next
			tail = prev
		}
		tail.Next = n2
		return n1
	}
	return n2
}

func SortedListMerge(n1 *NodeI, n2 *NodeI) *NodeI {
	MergLi(n1, n2)
	ListSort(n1)
	return n1
}

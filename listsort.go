package piscine

type NodeI struct {
	Data int
	Next *NodeI
}

func ListSort(l *NodeI) *NodeI {
	prev := l
	size := 0
	if prev != nil {
		for prev.Next != nil {
			size++
			prev = prev.Next
		}
		for i := 0; i <= size; i++ {
			prev = l
			for prev.Next != nil {
				if prev.Data > prev.Next.Data {
					prev.Data, prev.Next.Data = prev.Next.Data, prev.Data
				}
				prev = prev.Next
			}
		}
	}
	return l
}

package piscine

func BTreeApplyByLevel(root *TreeNode, f func(...interface{}) (int, error)) {
	if root == nil {
		return
	}
	var Qu []*TreeNode
	Qu = append(Qu, root)
	for len(Qu) > 0 {
		f(Qu[0].Data)
		if Qu[0].Left != nil {
			Qu = append(Qu, Qu[0].Left)
		}
		if Qu[0].Right != nil {
			Qu = append(Qu, Qu[0].Right)
		}
		Qu = Qu[1:]
	}
}

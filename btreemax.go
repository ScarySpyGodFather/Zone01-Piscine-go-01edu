package piscine

func BTreeMax(root *TreeNode) *TreeNode {
	current := root
	for current != nil {
		if current.Right != nil {
			current = current.Right
		} else {
			return current
		}
	}
	return current
}

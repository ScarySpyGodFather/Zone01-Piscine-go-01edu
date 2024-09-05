package piscine

func BTreeMin(root *TreeNode) *TreeNode {
	current := root
	for current != nil {
		if current.Left != nil {
			current = current.Left
		} else {
			return current
		}
	}
	return current
}

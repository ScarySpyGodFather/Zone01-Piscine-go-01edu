package piscine

func BTreeIsBinary(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if root.Left != nil && root != nil && root.Left.Data >= root.Data {
		return false
	} else if root.Right != nil && root != nil && root.Right.Data < root.Data {
		return false
	} else {
		if !BTreeIsBinary(root.Left) {
			return false
		}
		if !BTreeIsBinary(root.Right) {
			return false
		}
	}
	return true
}

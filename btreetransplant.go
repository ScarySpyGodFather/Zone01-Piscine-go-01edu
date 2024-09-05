package piscine

func BTreeTransplant(root, node, rplc *TreeNode) *TreeNode {
	if rplc.Data >= node.Parent.Data {
		node.Parent.Right = rplc
		rplc.Parent = node.Parent
		node.Parent = nil
	} else {
		node.Parent.Left = rplc
		rplc.Parent = node.Parent
		node.Parent = nil
	}
	return root
}

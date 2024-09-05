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


type TreeNode struct {
	Left, Right, Parent *TreeNode
	Data                string
}

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

func BTreeSearchItem(root *TreeNode, elem string) *TreeNode {
	current := root
	for current != nil {
		if elem == current.Data {
			return current
		}
		if current.Left != nil && elem < current.Data {
			current = current.Left
		} else if current.Right != nil && elem > current.Data {
			current = current.Right
		} else {
			return nil
		}
	}
	return nil
}

func BTreeInsertData(root *TreeNode, data string) *TreeNode {
	if root == nil {
		return &TreeNode{Data: data}
	}
	if data < root.Data {
		root.Left = BTreeInsertData(root.Left, data)
		root.Left.Parent = root

	} else {
		root.Right = BTreeInsertData(root.Right, data)
		root.Right.Parent = root
	}
	return root
}

func BTreeApplyInorder(root *TreeNode, f func(...interface{}) (int, error)) {
	if root == nil {
		return
	}
	BTreeApplyInorder(root.Left, f)
	f(root.Data)
	BTreeApplyInorder(root.Right, f)
}

func LastNode(root *TreeNode) *TreeNode {
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

func BTreeDeleteNode(root, node *TreeNode) *TreeNode {
	var forgotenBranche, nodeToBeReplaced *TreeNode
	if root == node {
		if root.Right != nil {
			forgotenBranche = root.Left
			root = root.Right
		} else {
			forgotenBranche = root.Right
			root = root.Left
		}
		nodeToBeReplaced = LastNode(root)
		node.Left = nil
		node.Right = nil
		node.Parent = nil
	} else {
		if node.Right != nil {
			forgotenBranche = node.Left
			nodeToBeReplaced = LastNode(node.Right)
			if node.Parent.Right == node {
				node.Parent.Right = node.Right
			} else if node.Parent.Left == node {
				node.Parent.Left = node.Right
			}
			node.Left = nil
			node.Right = nil
			node.Parent = nil
		} else if node.Left != nil {
			forgotenBranche = node.Right
			node.Right = nil
			node.Parent.Left = node.Left
			nodeToBeReplaced = LastNode(node.Right)
			if node.Parent.Right == node {
				node.Parent.Right = node.Right
			} else if node.Parent.Left == node {
				node.Parent.Left = node.Right
			}
			node.Left = nil
			node.Right = nil
			node.Parent = nil
		} else {
			if node.Parent.Right == node {
				node.Parent.Right = node.Right
			} else if node.Parent.Left == node {
				node.Parent.Left = node.Right
			}
			node.Parent = nil
		}
	}
	if nodeToBeReplaced != nil && forgotenBranche != nil {
		if forgotenBranche.Data >= nodeToBeReplaced.Data {
			BTreeInsertData(nodeToBeReplaced, string(int(nodeToBeReplaced.Data[0])+1))
			BTreeTransplant(root, nodeToBeReplaced.Right, forgotenBranche)
		} else {
			BTreeInsertData(nodeToBeReplaced, string(int(nodeToBeReplaced.Data[0])-1))
			BTreeTransplant(root, nodeToBeReplaced.Left, forgotenBranche)
		}
	}
	return root
}

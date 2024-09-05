package piscine

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

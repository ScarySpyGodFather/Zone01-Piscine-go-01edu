package piscine

func BTreeLevelCount(root *TreeNode) int {
	level := 0
	if root == nil {
		return level
	}
	levelL := BTreeLevelCount(root.Left)
	levelR := BTreeLevelCount(root.Right)
	if levelL <= levelR {
		level = levelR
	} else {
		level = levelL
	}
	level++
	return level
}

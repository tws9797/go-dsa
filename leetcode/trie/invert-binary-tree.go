package trie

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertTreeRecursively(root *TreeNode) *TreeNode {

	if root == nil {
		return nil
	}

	right := invertTreeRecursively(root.Right)
	left := invertTreeRecursively(root.Left)

	root.Right = left
	root.Left = right

	return root
}

package leetcode

func InvertTreeRecursive(root *TreeNode) *TreeNode {

	if root == nil {
		return nil
	}

	root.Left, root.Right = InvertTreeRecursive(root.Right), InvertTreeRecursive(root.Left)
	return root
}

func InvertTreeIteration(root *TreeNode) *TreeNode {

	var queue []*TreeNode
	var current *TreeNode
	if root == nil {
		return nil
	}
	queue = append(queue, root)
	for len(queue) > 0 {
		current = queue[0]
		queue = queue[1:]

		// Select one direction to enqueue
		current.Left, current.Right = current.Right, current.Left

		if current.Left != nil {
			queue = append(queue, current.Left)
		}

		if current.Right != nil {
			queue = append(queue, current.Right)
		}
	}

	return root
}

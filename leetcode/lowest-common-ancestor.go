package leetcode

import "go-dsa/leetcode/trie"

func LowestCommonAncestor(root, p, q *trie.TreeNode) *trie.TreeNode {

	currentNode := root

	for currentNode != nil {

		if p.Val > currentNode.Val && q.Val > currentNode.Val {
			currentNode = currentNode.Right
		} else if p.Val < currentNode.Val && q.Val < currentNode.Val {
			currentNode = currentNode.Left
		} else {
			return currentNode
		}
	}

	return nil
}

func LowestCommonAncestorRecursive(root, p, q *trie.TreeNode) *trie.TreeNode {

	currentNode := root

	for currentNode != nil {
		if p.Val > currentNode.Val && q.Val > currentNode.Val {
			return LowestCommonAncestorRecursive(currentNode.Right, p, q)
		} else if p.Val < currentNode.Val && q.Val < currentNode.Val {
			return LowestCommonAncestorRecursive(currentNode.Left, p, q)
		} else {
			return currentNode
		}
	}

	return nil
}

package leetcode

/*
Invert Binary Tree
Solution
Invert a binary tree.

Example:

Input:

     4
   /   \
  2     7
 / \   / \
1   3 6   9
Output:

     4
   /   \
  7     2
 / \   / \
9   6 3   1
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	return invert(root)
}

func invert(node *TreeNode) *TreeNode {
	if node.Left == nil && node.Right == nil {
		return node
	}

	var left *TreeNode
	if node.Left != nil {
		left = invert(node.Left)
	}
	var right *TreeNode
	if node.Right != nil {
		right = invert(node.Right)
	}
	node.Left = right
	node.Right = left
	return node
}

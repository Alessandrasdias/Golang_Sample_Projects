// The left subtree of a node contains only nodes with keys less than the node's key.
// The right subtree of a node contains only nodes with keys greater than the node's key.
// Both the left and right subtrees must also be binary search trees.

// The cases it won't be a binary search tree: Duplicate values or one level that doesn't satisfy the tree
// the pattern when going through a bst is that from the right side, the value the node must be greater than (root) persists, but the lower than updates
// the opposite occurs on the left side of the nodes
// so If I have a root of 15 and 20 goes on my right node, 19 as the left child node and 25 as the right child of 20,
// On the right, 25 has to be lower than positive infinity [which is awesome btw], but larger than 20
// On the left, 19 has to be lower then 20, but larger than 15
// On the left side of the root, the idea will follow almost the same way

// I think the a nice way to approach this is DFS - Pre order (Node, left, right)
package main

import "math"

//Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func dfs(node *TreeNode, min int, max int) bool {
	if node.Val <= min || node.Val >= max {
		return false
	}
	if node.Left != nil {
		if !dfs(node.Left, min, node.Val) {
			return false
		}
	}
	if node.Right != nil {
		if !dfs(node.Right, node.Val, max) {
			return false
		}
	}
	return true
}

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return dfs(root, math.MinInt32-1, math.MaxInt32+1)
}

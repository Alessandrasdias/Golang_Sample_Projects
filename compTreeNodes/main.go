// What is a complete tree? => Every level is completely full, but the last one don't need to be.
// However, at the last level, nodes must have be, as priority, at left first.
// For a tree to be full, every node needs to have 0 or 2 children, that means a node to the left could have 2 children, and on the same tree a node to the right could have 0 children and would still be valid
// The cases I could encounter in this situation would be: A complete full tree, a tree with only one node at the last level, and a tree with multiple possibilities for the last level (0 and 2)
// The values as not important because I just need to check if it's a valid tree, and return the number of the nodes in the tree not the values
// As it has to be a solution of less than 0(n), my options are o(log n), o(log^2 n)[log^2 being the hight of the tree] and o(1) which means Divide and conquer
// What can be divided?
// Well, If it's a complete full tree, than I can assume that everything before the last level is the "same" as in is with the same amount of doubling nodes, and everything on the last level can vary
// The last level can vary but must have it's nodes most to the left as possible, which means that if I find the most right node, I know that to the left nodes exist
// for the first part before the last level, I'll need to find the hight of the tree (without the last level)
// If I understand the last level as an array I can use binary search to find the index of the most right node
// When I'm doing the binary search I need to round up, otherwise I'll have the wrong value bc here I'm trying to find if the value(index) exists
package main

import "math"

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	height := treeHight(root)
	if height == 0 {
		return 1
	}
	upLevelCount := math.Pow(2, float64(height)) - 1
	left := 0
	right := upLevelCount
	for float64(left) < right {
		idx := math.Ceil((float64(left) + right) / 2)
		if nodeExists(int(idx), height, root) != nil {
			left = int(idx)
		} else {
			right = idx - 1
		}
	}
	return int(upLevelCount) + left + 1
}

func treeHight(root *TreeNode) int {
	height := 0
	for root.Left != nil {
		height++
		root = root.Left
	}
	return height
}

func nodeExists(idx int, height int, node *TreeNode) *TreeNode {
	left := 0.0
	right := math.Pow(2, float64(height)) - 1
	count := 0
	for count < height {
		midNode := math.Ceil((left + right) / 2)
		if idx >= int(midNode) {
			node = node.Right
			left = midNode
		} else {
			node = node.Left
			right = midNode - 1
		}
		count++
	}
	return node
}

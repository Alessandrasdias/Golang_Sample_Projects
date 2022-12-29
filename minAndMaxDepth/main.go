// bc the question is asking the max depth of the tree, dfs sounds like the best idea.
// DFS will go to the farthest node to the left first, however I can't be sure if the farthest node will be on the right or on the left
// bc of that i'll have to find a way of looking at both sides
// I have to keep moving to next nodes until I find the last one
// I'll use recursion bc there are several repetitive subtasks. left node and the Right node search
// When I hit a null point I'll stop
// Than it's just a matter of finding the Largest of the two (left and right)
// The values of the nodes don't matter in this question
package main

//Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func currMax(left, right int) int {
	if left > right {
		return left
	}
	return right
}
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return currMax(maxDepth(root.Left), maxDepth(root.Right)) + 1

}

// If I wanted to do the opposite and find the minimum depth the logic would stay partially the same
// What would change is the currMax that will compare to find lower numbers
// and some edge cases that have to be added such as:
// 1st -> If both, my left an my right nodes, while I'm at root, are nill, then I have to return 1 bc that how far it can go
// 2nd -> If my left is nil,check the minDepth of my right (same for right)
// If they are not nil then go about it as it would for the maxDepth

func currMin(left, right int) int {
	if left < right {
		return left
	}
	return right
}
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	} else if root.Left == nil && root.Right == nil {
		return 1
	} else {
		if root.Left == nil {
			return minDepth(root.Right) + 1
		}

		if root.Right == nil {
			return minDepth(root.Left) + 1
		}
	}

	return currMin(minDepth(root.Left), minDepth(root.Right)) + 1

}

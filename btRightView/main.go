// Because I'm looking from the right side some of the left nodes won't be seen, bc they'll be "overshadowed" by the right nodes on the same level
// If I figure out the right hight of the right side, I can have an idea of which nodes from left I can see, and which ones to ignore bc I won't see anyways
// So, my priority has to be the right most value first
// Also, I should keep track of the levels so I can have an idea of what to ignore
// DFS would be nice, but it's for the left, so I could try reversing the traversals for the right and see which of the results coming from them fits best with the potential final result I want
// The tree given by leetcode as an example is [1,2,3,null,5,null,4] the final answer should be [1,3,4]
// Reverse pre order would be [node, right, left], I would go all the way to the right and then left, knowing the levels it's just a matter of ignoring the correct left nodes (in this case all of them, but it would work if it had more on the left)
// Reverse pre order it's a better fit
// For the levels, I'll only save the values that are higher than the array length on the right, so only the left nodes  that don't have a right node on the same level will appear on the answer

package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
	result := make([]int, 0)
	dfs(root, 0, &result)
	return result

}

func dfs(node *TreeNode, curLevel int, result *[]int) {
	if node == nil {
		return
	}
	if curLevel >= len(*result) {
		*result = append(*result, node.Val)
	}
	if &node.Right != nil {
		dfs(node.Right, curLevel+1, result)
	}
	if &node.Left != nil {
		dfs(node.Left, curLevel+1, result)
	}
}

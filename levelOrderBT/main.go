// My return will have to be an array of arrays ex.: [[3],[6,1],[9,29],[5]]
// I'll need a queue to store the values and append those values to the result array at the end bc I have to hold the values for the levels and not just the values
// also to determine the level I need to know what is the first value of my queue, so I can expect the current values for the current level, but keep adding the children to the queue
// if the node has a left child push to the queue, same for the right side
// A level is when it finishes to process the current queue, so my current level is the length of the queue
// as far as I searched Golang does not provide an inbuilt queue structure
// I also need to keep track of my first value on the level
// If my root is null there is no point to continue

package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	result := [][]int{}
	queue := []*TreeNode{root}
	for len(queue) != 0 {
		curLevelValue := []int{} //subarray
		curLevel := len(queue)   // level
		for i := 0; i < curLevel; i++ {
			if queue[0] != nil {
				curLevelValue = append(curLevelValue, queue[0].Val)
				queue = append(queue, queue[0].Left)
				queue = append(queue, queue[0].Right)
			}
			queue = queue[1:]
		}
		result = append(result, curLevelValue)
	}
	return result[:len(result)-1]
}

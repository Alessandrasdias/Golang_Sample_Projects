// "Given an array of integers nums sorted in non-decreasing order" - Binary search
// so, the idea for BS is that O should have 2 pointer and a mid value that has to be left + right/2
// when left and right pointer overlap the value given is found (bc one will be at the beginning and the other at the most right)
// because I have no negative numbers I can return -1 if I don't find the value
// The pointer will move based on the comparison of mid and wanted value
// if wanted value is larger than mid value I have,  the left pointer will have to move 1 index after the mid to the right and ignore whatever was before

// for the solution I have to return starting and ending index bc the number will be repeated in the array
// Ex: [1,2,3,4,5,6,6,6,7,8]
// /.	0 1 2 3 4 5 6 7 8 9
// I need to return {5,7} when my target is 6
// to be able to recognize if the next number is still the target I'll have 4 variables: star, end, and have temp values for binary search within the subarrays on the left/right
// and I have to keep reducing the search space os the subarrays
// the first time I find the target may not be the first position
// So, when I use Binary search and it returns -1, I'll know that the index before/after has the first/last occurrence of the target

package main

//Logic for the Binary search
func bs(nums []int, left int, right int, target int) int {
	for left <= right {
		midValue := (left + right) / 2
		foundValue := nums[midValue]
		if foundValue == target {
			return midValue
		} else if foundValue < target {
			left = midValue + 1
		} else {
			right = midValue - 1
		}
	}
	return -1
}

//Solution finding the range of the target
func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1} // If I have an empty array return {-1, -1}
	}
	fstPos := bs(nums, 0, len(nums)-1, target) // search space is the entire array
	if fstPos == -1 {
		return []int{-1, -1} // If I can't find a match return {-1, -1}
	}
	start := fstPos // keep track of bs
	end := fstPos   // keep track of bs
	var temp1 int   // keep track of the last valid occurrence of target for start
	var temp2 int   // keep track of the last valid occurrence of target for the end
	for start != -1 {
		temp1 = start
		start = bs(nums, 0, start-1, target) // reducing search space to the left
	}
	start = temp1 // first occurrence of target
	for end != -1 {
		temp2 = end
		end = bs(nums, end+1, len(nums)-1, target) // reducing search space to the right
	}
	end = temp2 //last occurrence of target
	return []int{start, end}
}

//Space O(1)
//Time  O(log n)

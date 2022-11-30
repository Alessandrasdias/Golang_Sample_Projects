// is it sorted? no
// has negative numbers? yes
// brute force -> range through the array, identify the sub arrays, sum the sub arrays and get the largest sum 0(n^2)
// Divide the array to find out what's the sub array of that group and at the end compare groups?
// How to identify a subarray?
// array -> {-2, 1, -3, 4, -1, 2, 1, -5, 4}
// [-2][1][-3][4][-1][2][1][-5][4]
// [-2,1][-3,4][-1,2][1,-5][4]
// [-2,1,-3][4,-1,2][1,-5,4]
// [2,1,-3,4][-1,2][1,-5][4]
// [-2,-3] - not a subarray bc is not contiguous

// TODO brute force
// Identify the subarray and the sum of its values
// Save the sum of each
// identify the higher value among of subarray sum
// store result, compare, return highest

// !!! Kadane's Algorithm
// get one index and ask what's the maximum subarray ending in the current index? - reminded a bit of the logic of two sum, but only that will still be slow here
// faster if
// max subarray will be the current element or the current combined with the previous max subarray
// compare these to get the local max subarray, then it will be the global subarray to Compare to others

//0(n)
// one variable for current max
// one variable for the global max
// range through the array
// current max has to be current max + current element
// is current max higher then global max? no, is the current max less then 0? yes, current max reset to 0, keep looking
// is current max higher then global max? no, is the current max less then 0? no, keep looking
// is current max higher then global max? yes, so that is the new global max, return global max

// !! is the current max less then 0? -> bc if the array starts with a negative number throws a wrong output

// objective -> find the Maximum Subarray

package main

import "fmt"

func maxSubArray(nums []int) int {

	current_max := 0
	global_max := 0

	for i := range nums {
		current_max = current_max + nums[i]
		if current_max > global_max {
			global_max = current_max
		}
		if current_max < 0 {

			current_max = 0
		}
	}
	return global_max
}

func main() {

	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}

	fmt.Print(maxSubArray(nums))

}

//First of all I'll need to sort the array
//I'll quicksort the array in place. That means I won't break the array into smaller pieces.
//I'll use variables to represent the start and the end of the portions of the array I'll solve at each time
/* The logical process of quick sort/ partition func goes like this:
* I'll need a pivot
* The idea is to sort one element at a time based on that pivot (divide and conquer)
* i => will be a pointer to keep track of where is the real place of the pivot
* j => Will go through the array and compare the element at j to the pivot, if it's smaller swap j and i
* If it's larger don't do anything with i and move on with J
* My pivot needs to be at the right (most right at the end)
* and i and j most to the left
* When j value is equal to the pivot I need to swap, so pivot goes to i (rightful place)
* Now, I need to keep track of my partitions, because as quicksort goes, once I find the place of the pivot I still need to sort what ever is to the right and whatever is to the left of it
* So, If I had [5,3,2,4,6,1] -> [[1],2,[3,6,5,4]]
* let's imagine i put 2 in place, I need to understand that whatever is at left of two needs to be sorted, but the left pointer will be the same, and the right will be minus one(so, the position before two) so my range is correct
* and for the right, my left needs to be + one (so, the position after two), and my right remains the same
* Also, I'll use recursion and my base case is that my left should be smaller then my right
 */

/* I will have a swap function just to make it more readable, but it could be done without it. I want to use the swap because the logic will repeat itself twice
* the logic there is simple, it's just change positions of i and j, and i is stored in temp so I don't loose what as the original i before we start the swapping
 */

/*
*After the whole divide and conquer, sorting and partitioning game, I have to actually solve the problem at hand
*and for than I have the findKthLargest function
*My index to find will be whatever is the length of my array minus the k they are giving
* I'll call mu quicksort function and mark the pointers (left = 0 and Right = end of array)
* and return the result
 */

package main

func quickSort(array []int, left int, right int) {
	if left < right {
		var parIdx int
		parIdx = partition(array, left, right)
		quickSort(array, left, parIdx-1)
		quickSort(array, parIdx+1, right)

	}

}

func partition(array []int, left int, right int) int {
	pivot := array[right]
	parIdx := left
	for j := left; j < right; j++ {
		if array[j] < pivot {
			swap(array, parIdx, j)
			parIdx++
		}
	}
	swap(array, parIdx, right)
	return parIdx
}

func swap(array []int, i int, j int) {
	temp := array[i]
	array[i] = array[j]
	array[j] = temp
}

func findKthLargest(nums []int, k int) int {
	idxToFind := len(nums) - k
	quickSort(nums, 0, len(nums)-1)
	return nums[idxToFind]
}

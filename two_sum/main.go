package main

import "fmt"

//Given an array of integers
//return indices of the two numbers such that they add up to target
//each input would have exactly one solution, and you may not use the same element twice.
//return the answer in any order
//List is sorted
//no negative numbers

//-------------------
//pivot comparing which number on the list when added equals to target - Loop through entire list
// try to find is x is in the list:  2 - target(9) = x, but it's time consuming bc it has to go through the entire list every time it gets a pivot.

//How to make it better:
// Get the target and divide the list. Every number that is bigger than the target will not be compared.
// smaller sample to work with

//How to make even better:
// keeps checking if it has seen the number before (considering it's actual position)
//Ex.: "2 -> have I seen 7? No, ok, so let's try the next number and remember that at index of 0 I have num 2"
// " 7 - target(9) = x -> Have I seen 2 ? Yes, at index 0, so I'll return [0,1]".
// 2 e 7 cannot be used anymore.

// Using maps (hash table) we can get O(n)
// Actions needed:
// Create a variable to store the map that contains the number already seen
// loop to range through the array
// Create a variable to store the potential value that is equal to the target minus the number
// If the value found is equal to the potential value return the indexes
// If not, store the complement value and moves on

// recursive func and DP is also an option to solve it, it would be more readable and with the idea of storing on cache it would be o(n)
//It needs:
//cache to save index and num values
// A base case to make sure it won't run forever
// Call itself to repeat the subtasks

func twoSum(array []int, target int) []int {
	// Create a variable to store the map that contains the number already seen
	seenNumbs := map[int]int{}
	// loop to range through the array
	for i, num := range array {
		// Create a variable to store the potential value that is equal to the target minus the number
		potentialMatch := target - num
		// If the value found is equal to the potential value return the indexes
		if j, found := seenNumbs[potentialMatch]; found {
			return []int{j, i}

		}
		// If not, store the complement value and moves on
		seenNumbs[num] = i
	}
	return []int{}
}

func main() {

	// testing:

	nums := []int{2, 7, 11, 15}
	target := 9

	//nums := []int{3, 2, 4}
	//target := 6

	//nums := []int{3, 3}
	//target := 6

	fmt.Print(twoSum(nums, target))
}

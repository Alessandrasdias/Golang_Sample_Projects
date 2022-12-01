// sorted? no
// negative number? no
// decimal? no
//[1,2,4,5,6,2]

// func
// Create a map
// Range through the array
// Check if the num has been seen in the duplicateCheck before
// If the answer is yes. gets out of the loop and returns true
// If the answer is no keep looking
// o(n)

// Objective -> return true if there is a duplicate
package main

import "fmt"

func containsDuplicate(nums []int) bool {
	duplicateCheck := map[int]int{}
	for i, num := range nums {
		if _, found := duplicateCheck[num]; found {
			return true
		}

		duplicateCheck[num] = i
	}
	return false
}

func main() {

	//example := []int{1, 2, 3, 1}
	//example := []int{1, 2, 3, 1}
	//example := []int{1,2,3,4}
	example := []int{1, 2, 3, 1, 2, 2}

	fmt.Print(containsDuplicate(example))
}

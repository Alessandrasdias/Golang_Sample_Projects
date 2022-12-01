// is it sorted ? yes
// are there negatives? yes

// to rotate to the left:
// create a new slice
// loop through nums[] and while index of i is less then k value it will more forward and
// assign to my new array the number that is in the position of nums[1]
// and then I will append to my new array the position of num[0]
// nums gets my new array and then I can return it

package main

import "fmt"

func rotateLeft(nums []int, k int) []int {
	var newArray []int
	for i := 0; i < k; i++ {
		newArray = nums[1:]
		newArray = append(newArray, nums[0])
		nums = newArray
	}
	return nums
}

// To rotate to the right
// create a new slice
// I have one more param on this one, but it's no necessary and I will show that on the next solution
// loop through nums[] and while index of i is higher then k value and
// (i will start with the length of the array and will go backwards)
// assign to my new array the number that is in the position of nums[1]
// and then I will append to my new array the position of num[0]
// nums gets my new array and then I can return it

func rotateRight(nums []int, size int, k int) []int {
	var newArray []int
	for i := size; i > k; i-- {
		newArray = nums[1:]
		newArray = append(newArray, nums[0])
		nums = newArray
	}
	return nums
}

//same thing here but instead of having size, i just assign i to the length of nums

func rotateRight2(nums []int, k int) []int {
	var newArray []int
	for i := len(nums); i > k; i-- {
		newArray = nums[1:]
		newArray = append(newArray, nums[0])
		nums = newArray
	}
	return nums
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	nums2 := []int{-1, -100, 3, 99}

	fmt.Println("Rotating to the left:", rotateLeft(nums, 3))
	fmt.Println("Rotating to the left + negatives:", rotateLeft(nums2, 1))
	fmt.Println("Rotating to the left + negatives:", rotateLeft(nums2, 2))
	fmt.Println("Rotating to the Right with size: ", rotateRight(nums, 7, 3))
	fmt.Println("Rotating to the Right with size: + negatives:", rotateRight(nums2, 4, 2))
	fmt.Println("Rotating to the Right: ", rotateRight2(nums, 3))
	fmt.Println("Rotating to the Right + negatives: ", rotateRight2(nums2, 3))
}

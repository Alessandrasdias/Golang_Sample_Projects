// find the 0s in a array and move it to the end
// do not change the order of the non-zero numbers in the array
// cannot copy the array

// Only one loop
//j starting at 0
//i starting at 0
// range through the array
// if the value of my index i is different that 0 and my index i is different than J
// Swap -> what was at nums[i] will be the value of nums[j] , what was at nums[j] will be the value of nums[i]
// nums[i], nums[j] = nums[j], nums[i]

/* visualization:
The example: [0,1,0,3,12]
j=0
i=0
num[0] = 0  j=0
num[1] = 1  j=0 still bc will only change when  inside of the first if
|| here the swap ocurred bc j != i so -> (num[1] = 1), (nums[0] = 0)= (nums[0] = 1), (nums[1]=0)
|| [1,0]
gets out of the second if, j = 1 now loop continues
num[2] = 0  j=1  keep looping
|| [1,0,0]
num[3] = 3  j=1
|| here the swap ocurred bc j != i so -> (num[3] = 3), (nums[1] = 0) = (nums[1] = 3), (nums[3] = 0)
|| [1,0,0,3] -> [1,3,0,0]
gets out of the second if, j = 2 now loop continues
num[4] = 12 j=2
|| here the swap ocurred bc j != i -> (we are talking about indexes) so the value of j comes to the position of the value of i -> (num[4] = 12), (nums[2] = 0) = (nums[2] = 12), (nums[4] = 0)
|| [1,3,0,0,12] -> [1,3,12,0,0]

Because i will always be ahead position bc of 0s stopping j to move, it changes getting the zeros to the end of the array
*/

package main

import "fmt"

func moveZeroes(nums []int) []int {
	j := 0
	for i := range nums {
		if nums[i] != 0 {
			if i != j {
				nums[i], nums[j] = nums[j], nums[i]
			}
			j++
		}
	}

	return nums

}

func main() {

	nums := []int{0, 1, 0, 3, 12}
	//nums := []int{0, 1, 0, 3, 0}
	//nums := []int{0, 1, 0, 3, 12, 9, 0, 10}

	fmt.Print(moveZeroes(nums))

}

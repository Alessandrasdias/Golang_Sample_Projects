package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'miniMaxSum' function below.
 *
 * The function accepts INTEGER_ARRAY arr as parameter.
 */
// example array 1 2 3 4 5 and 7 69 2 221 8974
// I need a variable for min and max that will start with the value of the index 0
// a variable to hold the total_sum that will start at 0
// the idea will be that the total_sum, at the end will be the sum of all indexes
// so i'll loop through the array
// curr_sum will be the value i have (0 at first) plus the value of the index in question
//if my min (that is now 1) is lower then the value of my index at the moment (is 1)
// then my min is equal to my index value
//if my max (that is now 1) is higher then the value of my index at the moment (is 1)
// then my max is equal to my index value
// at the first time it will me equal because min and max already starts with the value of index 0
// ---for index 1
//total_sum = 3
// min = 1 is < 3 -> yes min= 2
// max = 1 > 3 -> no max = 1
// ---for index 2
// total_sum = 3 + 3 = 6
// min = 3 is < 3 -> yes min= 3
// max = 1 > 6 -> no max = 1
// ---for index 3
// total_sum = 6 + 4 = 10
// min = 6 is < 10 -> yes min= 4
// max = 1 > 10 -> no max = 1
// ---for index 4
// total_sum = 10 + 5 = 15
// min = 10 is < 15 -> yes min= 5
// max = 1 > 15 -> no max = 1

// now I have to get my total_sum at subtract the values of min and max
// total_sum(15) - min(5)
// total_sum(15) - max(1)
// min will be 10
// max will be 14

// what happens is, whatever that was on [0] has to be excluded from the sum of the max
// and whatever is at [4] has to be excluded from sum of the min

// Also, it's important to have it converted to int64 bc we could have very large numbers

func miniMaxSum(arr []int32) {
	var min int64 = (int64(arr[0]))
	var max int64 = (int64(arr[0]))
	var total_sum int64 = 0

	for i := range arr {
		total_sum = total_sum + int64(arr[i])
		if min < int64(arr[i]) {
			min = int64(arr[i])
		}
		if max > int64(arr[i]) {
			max = int64(arr[i])
		}
	}

	fmt.Printf("%d\t", total_sum-min)
	fmt.Printf("%d", total_sum-max)
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	arrTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var arr []int32

	for i := 0; i < 5; i++ {
		arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
		checkError(err)
		arrItem := int32(arrItemTemp)
		arr = append(arr, arrItem)
	}

	miniMaxSum(arr)
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

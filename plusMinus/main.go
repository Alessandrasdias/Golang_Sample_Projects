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
 * Complete the 'plusMinus' function below.
 *
 * The function accepts INTEGER_ARRAY arr as parameter.
 */
//Print the ratios of positive, negative and zero values in the array.
//Each value should be printed on a separate line with 6 digits after the decimal.
//The function should not return a value.

/*
TODO

* Whatever I have it will have to be divided by the number of elements of the array, Which means at some point I'll have ratio/n
* I can have negatives, zeros and positives
* Is it sorted? If yes, I can count how many numbers are there before and after zero and then divide the ratio by the size of the array
* Assuming they are not sorted
* I can have a counter and every time I see a neg number I add - same for positive
* Zero is not considered pos or neg, so if it finds a zero add to a third counter

* Let's range through the array
* 1st are you zero ?
* 2nd are you lower than 0?
* add positive
*/
var pos float32
var neg float32
var zero float32

func plusMinus(arr []int32) {

	for _, num := range arr {
		if num == 0 {
			zero++
		}
		if num > 0 {
			pos++
		}
		if num < 0 {
			neg++
		}
	}
	/*Testings bc better safe than sorry =P
	  fmt.Println(pos, neg, zero)*/

	resPos := pos / float32((len(arr)))
	resNeg := neg / float32((len(arr)))
	resZero := zero / float32((len(arr)))

	fmt.Println(resPos)
	fmt.Println(resNeg)
	fmt.Println(resZero)

}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	n := int32(nTemp)

	arrTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var arr []int32

	for i := 0; i < int(n); i++ {
		arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
		checkError(err)
		arrItem := int32(arrItemTemp)
		arr = append(arr, arrItem)
	}

	plusMinus(arr)

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

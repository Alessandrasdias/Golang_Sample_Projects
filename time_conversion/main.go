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
 * Complete the 'timeConversion' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts STRING s as parameter.
 */

// minutes and seconds don't really have to be altered, they will be the same in both situations
// the goal is to change the numbers from after 12 moon to midnight bc the rest is the same
// that means when I have 1pm i have to show 13h
// I'll have to split the part of the string that is the hour, change it and return the entire string (i don't need to change the minutes and seconds, but I have to return them )
// also I'll have to return without the pm/am part
// Before anything I need to select only those hours that are PM, bc AM won't change anyways
// unless it's 12am, then it's the only case I have to change for 00 when it's AM

// check if it's am or pm
// have a slice minus the suffix
// split array according to separator
// assign hh,mm, ss to a new array
// convert the hour to int so we can check
//then convert back to string and return result

func timeConversion(s string) string {
	pm := strings.HasSuffix(s, "PM")
	am := strings.HasSuffix(s, "AM")

	time := s[:len(s)-2]

	timeArr := strings.Split(time, ":")
	hh, mm, ss := timeArr[0], timeArr[1], timeArr[2]
	hhInt, err := strconv.Atoi(hh)

	if err != nil {
		panic(err.Error())
	}

	if am && hhInt == 12 {
		hhInt = 0
	}

	if pm && hhInt != 12 {
		hhInt += 12
	}

	hh = strconv.Itoa(hhInt)

	return fmt.Sprintf("%02s:%02s:%02s", hh, mm, ss)
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	s := readLine(reader)

	result := timeConversion(s)

	fmt.Fprintf(writer, "%s\n", result)

	writer.Flush()
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

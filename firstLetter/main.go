package main

import "fmt"

// lowercase English letters
// TODO to prevent other type letters:
// Convert letter to lowercase no matter the input
// If another language alphabet letters are entered the program would have to understand it doesn't belong to english and treat the error
// bc the character would be a different value and the program would return the wrong result

//s will contain at least one letter that appears twice - no case of no letter repeated
// but if that was an issue the error would have to be handled and display something like:
// "The string entered does not contain repeated letter"

// It's not about every repetition, but the first repetition found

//Solving the problem:
// If I try to search through the entire list for every character and compare it would take too long and require nested for loops
// O(n^2)

// I could range through the string and at each character check if I had seen a character like that before
// O(n)

// func repeatedCharacter
// Create a map
// Create a variable to store the number of repetitions
// Range through the string
// Check if the char has been seen in the charSet before
// If the answerer is yes, stores the number of time it has been seen, gets out of the loop and returns the char
// If the answerer is no keep looking
// o(n)

func repeatedCharacter(s string) byte {
	charSet := map[rune]int{}
	rep := 0
	for i, char := range s {
		if _, found := charSet[char]; found {
			rep = i
			break
		}
		charSet[char] = i
	}

	return s[rep]

}

func main() {

	ch := repeatedCharacter("abccbaacz")
	fmt.Println(string(ch)) // because of how Go works teh answerer needs to be converted

}

// use trie (prefix trie)
// null root
// o(length of the world)
// Go has a package for tries =>  go get github.com/dghubble/trie
// https://pkg.go.dev/github.com/dghubble/trie#section-readme
// needs dictionary
// tries are much more efficient for this type of situation though

// Op 2, takes longer and more space
// Create a pre allocated map for parts where the key is a string and the value is a struct that takes the length of parts
// range through parts and add the parts to a struct
// range through the words
// result takes the value of the word
// create a loop for the word where if there are any parts that exists inside the partmap
// break the loop and add to results

package main

import "fmt"

func solution(words []string, parts []string) []string {
	partsMap := make(map[string]struct{}, len(parts))
	for _, part := range parts {
		partsMap[part] = struct{}{}
	}
	results := make([]string, 0, len(words))
	for _, word := range words {
		result := word
	wordLoop:
		for j := 5; j >= 1; j-- {
			for i := 0; i+j <= len(word); i++ {
				if _, exists := partsMap[word[i:i+j]]; exists {
					result = word[0:i] + "[" + word[i:i+j] + "]" + word[i+j:]
					break wordLoop
				}
			}
		}
		results = append(results, result)
	}
	return results
}

func main() {
	words := []string{"Apple", "Melon", "Orange", "Watermelon"}
	parts := []string{"a", "mel", "lon", "el", "An"}
	fmt.Print(solution(words, parts))
}

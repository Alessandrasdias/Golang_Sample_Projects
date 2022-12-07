//TODO
// 1st create a global slice with the possible strings

var table = []string{"", "", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}

func solution(buttons string) []string {

	var result []string
	helper(&result, buttons, "", 0)
	return result

}

// 2nd  create a helper function where:
// the result will be a pointer to the slice of strings
// I'll have the buttons and prefix as a string
// and an index
func helper(result *[]string, buttons, prefix string, index int) {
	if index >= len(buttons) { //if my index is higher or equal to the length of buttons
		if len(prefix) > 0 { // and the length of my prefix is higher than 0
			*result = append(*result, prefix) // add to my result the prefix
		}
		return
	}
	digit := buttons[index] - '0' //  digit is the buttons index
	if digit < 2 {                // if my digit is lower than 2
		helper(result, buttons, prefix, index+1) // call my helper func, move to the next index
	} else {
		for _, v := range table[digit] { // range through the table of digits
			helper(result, buttons, prefix+string(v), index+1)
			// call my func helper. Get the the prefix + the value of the string at the index v

		}
	}
}

// alternative way

/*
func letterCombinations(digits string) []string {
    result := make([]string,0)
    charMapping := []string{"0","1","abc","def","ghi","jkl","mno","pqrs","tuv","wxyz"}

    if(len(digits) != 0){
        result = append(result, "")

        for i := 0; i < len(digits); i++ {
            index, _ := strconv.Atoi(string(digits[i]))
            for len(result[0]) == i {
                permutation := result[0]
                result = result[1:]
                for _, c := range charMapping[index] {
                    result = append(result, permutation+string(c))
                }
            }
        }
    }

    return result
}
*/
// 1st create a way to validate is there are no duplicates for rows or columns (func to validate if there are duplicates in a regular array - use hashmap )
// 2nd for the solution, first check if there are any duplicates in the grid rows bc if yes it's already invalid - game rule
// 3rd range through rows and columns of the grid
// 4th assign the rows and columns of the grid to a new board
// 5th check for duplicates on the board
// for the subgrid 3x3 the logic is similar
// but first we have to find i, j in the grid that are = 3 - game rule
// then we create a slice for the grid and do the same we did with the board

// not the best solution since we have nested for loops
// note to self -> think about better answer later
package main

func hasDuplicates(a []string) bool {
	seen := map[string]bool{}
	for _, c := range a { //ranges through the array
		if c == "." { // if it finds a dot continues
			continue
		}
		if seen[c] { // is I had seen it before return true
			return true
		}
		seen[c] = true // remember I have seen it to compare next time
	}
	return false
}

func solution(grid [][]string) bool {
	for i := 0; i < len(grid); i++ { // while my i is lower than the length
		if hasDuplicates(grid[i]) { // check for duplicates in my grid row
			return false // return false if you find them
		}
	}
	for col := 0; col < len(grid); col++ { //while my col is lower than the length of the grid
		board := []string{}                    // my board is a a slice of strings
		for row := 0; row < len(grid); row++ { //while my row is lower than the length of grid
			board = append(board, grid[row][col]) // I need to append to my board the values for row and column
		}
		if hasDuplicates(board) { // then check for duplicates rows and columns on my board
			return false
		}
	}
	//subGrid 3x3
	//while my col i is lower than the length of the grid (keep going until is i == 3)
	for i := 0; i < len(grid); i += 3 {
		for j := 0; j < len(grid); j += 3 {
			subGrid := []string{} //my subGrid is a slice of strings
			for t := 0; t < 3; t++ {
				for s := 0; s < 3; s++ {
					subGrid = append(subGrid, grid[i+t][s+j]) // I need to append to my subGrid the values for row and column
				}
			}
			if hasDuplicates(subGrid) { // then check for duplicates
				return false
			}
		}
	}
	return true
}

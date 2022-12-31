// One of the most important things here is to understand the definition of island and what is not an island
// For this activity, an island is a land(1) connected horizontally or vertically
// if the land (1) is alone and diagonal to the others, they are separate islands
/*
	["0","1","0","1","0"]
  	["1","0","1","0","0"]
  	["0","1","1","1","0"]
  	["1","0","0","0","0"]

There would be 5 islands here

*/
// Whenever I find a 1 (land), it could be a new land or a part of a land and then could happen a double count
// To prevent that, If I find a 1 I must ignore its adjacent 1s
// As I'm already ignoring 0s, I could turn adjacent 1s to 0
// I will need to go through every element do matter the order and while I'm doing that I need to shift all the 1s
// I'll use dfs
//My dfs function will take the grid, my current column and row to switch 1s, but not before validating the grid, that's all it'll do
// My numIslands is looping through the grid and looking for ones and adding to island

package main

func dfs(grid [][]byte, curRow int, curCol int) {

	if curRow < 0 || curRow >= len(grid) || curCol < 0 || curCol >= len((grid)[0]) || grid[curRow][curCol] == '0' {
		return
	}
	grid[curRow][curCol] = '0'
	dfs(grid, curRow+1, curCol)
	dfs(grid, curRow-1, curCol)
	dfs(grid, curRow, curCol+1)
	dfs(grid, curRow, curCol-1)

}

func numIslands(grid [][]byte) int {
	island := 0
	for row := range grid {
		for col := range grid[0] {
			if grid[row][col] == '1' {
				dfs(grid, row, col)
				island += 1
			}
		}
	}
	return island
}

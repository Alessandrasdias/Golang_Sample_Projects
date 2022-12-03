// note to self: think about a better solution
package main

import "fmt"

func solution(a [][]int) [][]int {

	for i, temp, n := 0, 0, len(a)-1; i <= n/2; i++ {
		for j := i; j < n-i; j++ {
			temp = a[j][n-i]
			a[j][n-i] = a[i][j]
			a[i][j] = a[n-j][i]
			a[n-j][i] = a[n-i][n-j]
			a[n-i][n-j] = temp
		}
	}
	return a
}

func main() {
	matrix := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}

	fmt.Println(solution(matrix))
}

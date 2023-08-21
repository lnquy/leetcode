package main

import "fmt"

func main() {
	fmt.Println(flipAndInvertImage([][]int{{1, 1, 0}, {1, 0, 1}, {0, 0, 0}}))                        // [[1,0,0],[0,1,0],[1,1,1]]
	fmt.Println(flipAndInvertImage([][]int{{1, 1, 0, 0}, {1, 0, 0, 1}, {0, 1, 1, 1}, {1, 0, 1, 0}})) // [[1,1,0,0],[0,1,1,0],[0,0,0,1],[1,0,1,0]]
}

func flipAndInvertImage(a [][]int) [][]int {
	for i := range a {
		n := len(a[i])
		if n%2 == 1 {
			a[i][n/2] = 1 ^ a[i][n/2]
		}
		for j := 0; j < n/2; j++ {
			a[i][j], a[i][n-1-j] = 1^a[i][n-1-j], 1^a[i][j]
		}
	}
	return a
}

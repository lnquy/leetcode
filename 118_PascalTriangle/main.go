package main

import "fmt"

func main() {
	fmt.Println(generate(5))
	//fmt.Println(generate_big(5))
}

func generate(numRows int) [][]int {
	out := make([][]int, numRows, numRows)
	if numRows < 1 {
		return out
	}

	out[0] = []int{1}
	for i := 1; i < numRows; i++ {
		out[i] = make([]int, i+1, i+1)
		out[i][0], out[i][i] = 1, 1

		for j := i - 1; j > 0; j-- {
			out[i][j] = out[i-1][j] + out[i-1][j-1]
		}
	}
	return out
}

// This solution run faster on big N (>100) but slower on small N since it has to do more bound checking.
func generate_big(numRows int) [][]int {
	out := make([][]int, numRows, numRows)
	if numRows < 1 {
		return out
	}

	out[0] = []int{1}
	for i := 1; i < numRows; i++ {
		out[i] = make([]int, i+1, i+1)
		out[i][0], out[i][i] = 1, 1

		for j := 1; j < (i+1)/2; j++ {
			s := out[i-1][j] + out[i-1][j-1]
			out[i][j], out[i][i-j] = s, s
		}
		if i > 1 && (i+1)%2 == 1 {
			out[i][i/2] = out[i-1][i/2] + out[i-1][i/2-1]
		}
	}
	return out
}

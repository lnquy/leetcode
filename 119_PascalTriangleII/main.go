package main

import "fmt"

func main() {
	fmt.Println(getRow(3))
}

func getRow(rowIndex int) []int {
	out := make([]int, rowIndex+1, rowIndex+1)
	out[0] = 1
	if rowIndex == 0 {
		return out
	}

	for i := 1; i <= rowIndex; i++ {
		out[0], out[i] = 1, 1
		carry, prev := 0, 1
		for j := 1; j <= i; j++ {
			if carry > 0 {
				prev = out[j-1]
				out[j-1] = carry
			}
			carry = out[j] + prev
		}
	}
	return out
}

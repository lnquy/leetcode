package main

import "fmt"

func main() {
	fmt.Println(majorityElement([]int{3, 2, 3}))
	fmt.Println(majorityElement([]int{2, 2, 1, 1, 1, 2, 2}))
}

// Boyer-Moore Voting Algorithm
func majorityElement(nums []int) int {
	out, c := 0, 0

	for _, n := range nums {
		if c == 0 {
			out = n
		}
		if out == n {
			c++
		} else {
			c--
		}
	}
	return out
}

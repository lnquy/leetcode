package main

import "fmt"

func main() {
	fmt.Println(removeElement([]int{3, 2, 2, 3}, 3))
	fmt.Println(removeElement([]int{0, 1, 2, 2, 3, 0, 4, 2}, 2))
}

func removeElement(nums []int, val int) int {
	n := len(nums)

	for i := 0; i < n; i++ {
		if nums[i] == val {
			nums[i], nums[n-1] = nums[n-1], nums[i]
			n--
			i--
		}
	}
	return n
}

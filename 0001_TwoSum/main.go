package main

import "fmt"

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
}

// O(n)
func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, num := range nums {
		if idx, ok := m[target-num]; ok {
			return []int{idx, i}
		}
		m[num] = i
	}
	return nil
}

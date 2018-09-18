package main

import "fmt"

func main() {
	fmt.Println(removeDuplicates([]int{1, 1, 2}))
	fmt.Println(removeDuplicates([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}))
}

// Idea: If duplicated then swap current item to nums[l-1] with `l = len(nums) - #duplicated`
// This solution will return [0,4,1,3,2] list instead of [0,1,2,3,4] but has better time complexity O(n-#dup)
// and not require temporary array as slicing solution.
func removeDuplicates_OriginalSolution(nums []int) int {
	m := map[int]struct{}{}

	l := len(nums)
	for i := 0; i < l; i++ {
		if _, ok := m[nums[i]]; ok {
			nums[i], nums[l-1] = nums[l-1], nums[i]
			l--
			i--
			continue
		}
		m[nums[i]] = struct{}{}
	}
	return l
}

func removeDuplicates(nums []int) int {
	m := map[int]struct{}{}

	l := len(nums)
	for i := 0; i < l; i++ {
		if _, ok := m[nums[i]]; ok {
			nums = append(nums[:i], nums[i+1:]...)
			l--
			i--
			continue
		}
		m[nums[i]] = struct{}{}
	}
	return l
}

func removeDuplicates_LeetCodeSolution(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	i := 0
	for j := 1; j < len(nums); j++ {
		if nums[j] != nums[i] {
			i++
			nums[i] = nums[j]
		}
	}
	return i + 1
}

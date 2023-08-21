package main

func moveZeroes(nums []int) {
	nonZeroIndex := -1
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == 0 {
			if nonZeroIndex < i {
				nonZeroIndex = i
			}
			nonZeroIndex = firstNonZero(nums, nonZeroIndex)
			if nonZeroIndex == -1 {
				break
			}
			nums[i], nums[nonZeroIndex] = nums[nonZeroIndex], nums[i] // Swap 0 and non-0
		}
	}
}

func firstNonZero(nums []int, startIdx int) int {
	for i := startIdx; i < len(nums); i++ {
		if nums[i] != 0 {
			return i
		}
	}
	return -1
}

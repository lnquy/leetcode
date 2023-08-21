package main

import "math"

func increasingTriplet(nums []int) bool {
	min := nums[0]
	secondMin := math.MaxInt

	for i := 1; i < len(nums); i++ {
		n := nums[i]
		if n <= min {
			min = n
		} else if n <= secondMin {
			secondMin = n
		} else {
			return true
		}
	}
	return false
}

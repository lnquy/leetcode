package main

func findMaxAverage(nums []int, k int) (maxAvg float64) {
	lastSum := 0

	for i := 0; i < k; i++ {
		lastSum += nums[i]
	}
	maxAvg = float64(lastSum) / float64(k)
	if len(nums) == k {
		return maxAvg
	}

	for i := k; i < len(nums); i++ {
		lastSum = lastSum - nums[i-k] + nums[i]
		avg := float64(lastSum) / float64(k)
		if avg > maxAvg {
			maxAvg = avg
		}
	}

	return maxAvg
}

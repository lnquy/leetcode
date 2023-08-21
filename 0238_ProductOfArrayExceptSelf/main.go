package main

func productExceptSelf(nums []int) (result []int) {
	numLen := len(nums)
	left, result := make([]int, numLen, numLen), make([]int, numLen, numLen)
	left[0] = 1
	// Calculate product of all items on the left of num[i]
	for i := 1; i < numLen; i++ {
		left[i] = left[i-1] * nums[i-1]
	}
	// Calculate product of all items on the right of num[i]
	// Then also calculate the result
	lastRight := 1
	for i := numLen - 2; i >= 0; i-- {
		lastRight *= nums[i+1]
		result[i] = left[i] * lastRight
	}
	// Calculate the last result item, as on the right for loop we ignored the last index
	result[numLen-1] = left[numLen-1]

	return result
}

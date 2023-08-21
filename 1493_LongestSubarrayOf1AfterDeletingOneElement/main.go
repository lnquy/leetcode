package main

func longestSubarray(nums []int) (longest int) {
	l, r, zeroes := 0, 0, 0

	for ; r < len(nums); r++ {
		zeroes += 1 - nums[r]

		if zeroes > 1 {
			zeroes -= 1 - nums[l]
			l++
		}
	}

	return r - l - 1
}

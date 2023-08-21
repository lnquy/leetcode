package main

func longestOnes(nums []int, k int) (longest int) {
	l, r, currZeros := 0, 0, 0

	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			currZeros++
		}
		for currZeros > k {
			if nums[l] == 0 {
				currZeros--
			}
			l++
		}

		r++
		if longest < r-l {
			longest = r - l
		}
	}
	return longest
}

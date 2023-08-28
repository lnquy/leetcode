package main

func findPeakElement(nums []int) (mid int) {
	l := len(nums)
	if l == 1 || nums[0] > nums[1] { // First item is peak
		return 0
	}
	if nums[l-1] > nums[l-2] { // Last item is peak
		return l - 1
	}

	left, right := 1, l-2
	for left <= right {
		mid = left + (right-left)/2
		if nums[mid] > nums[mid-1] && nums[mid] > nums[mid+1] { // Peak
			return mid
		}
		if nums[mid] < nums[mid+1] { // Continue to search on the right potion
			left = mid + 1
			continue
		}
		if nums[mid] < nums[mid-1] { // Continue to search on the left potion
			right = mid - 1
			continue
		}
	}
	return mid
}

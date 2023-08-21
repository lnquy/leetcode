package main

func maxArea(height []int) int {
	l, r, area := 0, len(height)-1, 0

	for l < r {
		a := min(height[l], height[r]) * (r - l)
		if a > area {
			area = a
		}
		if height[l] > height[r] {
			r--
		} else {
			l++
		}
	}
	return area
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

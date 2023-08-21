package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(findMedianSortedArrays([]int{1, 3}, []int{2}))
	fmt.Println(findMedianSortedArrays([]int{1, 2}, []int{3, 4}))
}

// O(log(m+n))
func findMedianSortedArrays(nums1, nums2 []int) float64 {
	m, n := len(nums1), len(nums2)
	if m > n {
		m, n = n, m
		nums1, nums2 = nums2, nums1
	}
	iMin, iMax, mid := 0, m, (m+n+1)/2
	for iMin <= iMax {
		i := (iMin + iMax) / 2
		j := mid - i
		if i < iMax && nums2[j-1] > nums1[i] {
			iMin = i + 1
		} else if i > iMin && nums1[i-1] > nums2[j] {
			iMax = i - 1
		} else {
			maxLeft := 0
			if i == 0 {
				maxLeft = nums2[j-1]
			} else if j == 0 {
				maxLeft = nums1[i-1]
			} else {
				maxLeft = max(nums1[i-1], nums2[j-1])
			}
			if (m+n)%2 == 1 {
				return float64(maxLeft)
			}

			minRight := 0
			if i == m {
				minRight = nums2[j]
			} else if j == n {
				minRight = nums1[i]
			} else {
				minRight = min(nums2[j], nums1[i])
			}

			return float64(maxLeft+minRight) / 2.0
		}
	}
	return 0.0
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// O((m+n)/2)
func findMedianSortedArrays_Slow(nums1, nums2 []int) float64 {
	m, n := len(nums1), len(nums2)
	s := ((m + n) / 2) + 1

	out := make([]int, 0, s)
	a, b, n1, n2 := 0, 0, math.MaxInt64, math.MaxInt64
	for i := 0; i < s; i++ {
		if a >= m {
			n1 = math.MaxInt64
		} else {
			n1 = nums1[a]
		}
		if b >= n {
			n2 = math.MaxInt64
		} else {
			n2 = nums2[b]
		}

		if n1 < n2 {
			out = append(out, n1)
			a++
		} else {
			out = append(out, n2)
			b++
		}
	}

	if (m+n)%2 == 0 {
		return float64(out[s-2]+out[s-1]) / 2.0
	}
	return float64(out[s-1])
}

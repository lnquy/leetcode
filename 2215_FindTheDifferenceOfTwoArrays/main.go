package main

func findDifference(nums1 []int, nums2 []int) [][]int {
	m1 := make(map[int]struct{}, len(nums1))
	for _, n := range nums1 {
		m1[n] = struct{}{}
	}

	m2 := make(map[int]struct{}, len(nums2))
	for _, n := range nums2 {
		m2[n] = struct{}{}
	}

	result := make([][]int, 2, 2)
	for k := range m1 {
		if _, ok := m2[k]; !ok {
			result[0] = append(result[0], k)
		}
	}
	for k, _ := range m2 {
		if _, ok := m1[k]; !ok {
			result[1] = append(result[1], k)
		}
	}

	return result
}

package main

func uniqueOccurrences(arr []int) bool {
	m := make(map[int]int, len(arr))
	for _, n := range arr {
		m[n]++
	}

	m2 := make(map[int]struct{}, len(m))
	for _, v1 := range m {
		if _, ok := m2[v1]; ok {
			return false
		}
		m2[v1] = struct{}{}
	}
	return true
}

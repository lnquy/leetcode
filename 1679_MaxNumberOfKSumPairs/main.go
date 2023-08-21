package main

func maxOperations(nums []int, k int) (op int) {
	m := make(map[int]int, len(nums))

	for _, n := range nums {
		delta := k - n
		if counter, ok := m[delta]; ok {
			if counter == 1 {
				delete(m, delta)
			} else if counter > 1 {
				m[delta] = counter - 1
			}
			op++
			continue
		}

		if counter, ok := m[n]; ok {
			m[n] = counter + 1
		} else {
			m[n] = 1
		}
	}

	return op
}

package main

func dailyTemperatures(temperatures []int) []int {
	res := make([]int, len(temperatures), len(temperatures))
	stack := make([]int, 0, len(temperatures))

	for i, todayTemp := range temperatures {
		for len(stack) > 0 && todayTemp > temperatures[stack[len(stack)-1]] {
			lastStackIdx := len(stack) - 1
			res[stack[lastStackIdx]] = i - stack[lastStackIdx]
			stack = stack[:len(stack)-1]
		}

		stack = append(stack, i)
	}

	return res
}

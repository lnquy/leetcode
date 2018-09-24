package main

import "fmt"

func main() {
	fmt.Println(isMonotonic([]int{1, 2, 2, 3}))
	fmt.Println(isMonotonic([]int{6, 5, 4, 4}))
	fmt.Println(isMonotonic([]int{1, 3, 2}))
	fmt.Println(isMonotonic([]int{1, 2, 4, 5}))
	fmt.Println(isMonotonic([]int{1, 1, 1}))
}

func isMonotonic(a []int) bool {
	n := len(a)
	if n == 1 {
		return true
	}

	sign := 0
	for i := 1; i < n; i++ {
		diff := a[i] - a[i-1]
		if sign == 0 {
			if diff != 0 {
				sign = diff
			}
			continue
		}
		if (sign > 0 && diff < 0) || (sign < 0 && diff > 0) {
			return false
		}
	}
	return true
}

package main

import "fmt"

func main() {
	fmt.Println(distributeCandies([]int{1, 1, 2, 2, 3, 3}))
	fmt.Println(distributeCandies([]int{1, 1, 2, 3}))
}

func distributeCandies(candies []int) int {
	m := make(map[int]struct{})

	for _, c := range candies {
		if _, ok := m[c]; !ok {
			m[c] = struct{}{}
		}
	}
	if len(m) < len(candies)/2 {
		return len(m)
	}
	return len(candies) / 2
}

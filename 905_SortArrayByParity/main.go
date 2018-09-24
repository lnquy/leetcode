package main

import "fmt"

func main() {
	fmt.Println(sortArrayByParity([]int{3, 1, 2, 4}))
}

func sortArrayByParity(a []int) []int {
	n := len(a)
	for i := 0; i < n; i++ {
		if a[i]%2 == 1 {
			a[i], a[n-1] = a[n-1], a[i]
			n--
			i--
		}
	}
	return a
}

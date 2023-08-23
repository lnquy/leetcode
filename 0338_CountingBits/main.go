package main

func countBits(n int) []int {
	res := make([]int, n+1)
	for i := range res {
		res[i] = res[i>>1] + i&1
	}
	return res
}

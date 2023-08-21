package main

func main() {
	pivotIndex([]int{1, 7, 3, 6, 5, 6})
}

func pivotIndex(nums []int) int {
	sum := 0
	for _, n := range nums {
		sum += n
	}

	sumOnLeftSide := 0
	for i, n := range nums {
		sum -= n
		if sumOnLeftSide == sum {
			return i
		}
		sumOnLeftSide += n
	}
	return -1
}

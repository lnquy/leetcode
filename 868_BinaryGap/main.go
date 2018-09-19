package main

import "fmt"

func main() {
	fmt.Println(binaryGap(22))
	fmt.Println(binaryGap(5))
	fmt.Println(binaryGap(6))
	fmt.Println(binaryGap(8))
}

// O(log(n))
func binaryGap(n int) int {
	gap, last1BitIdx := 0, -1

	i := 0
	for n != 0 {
		if n & 1 != 0 {
			if last1BitIdx >= 0 && i - last1BitIdx > gap {
				gap = i - last1BitIdx
			}
			last1BitIdx = i
		}

		n = n >> 1
		i++
	}
	return gap
}

// This function is the solution for Codility's Binary Gap exercise (https://app.codility.com/demo/results/trainingECGAX7-EVD/#),
// not LeetCode exercise.
// It asks for finding maximum consecutive 0 gap bits. These two problems are directly related so I put it here, also.
// Another solution for this exercise is using result from LeetCode's solution - 1.
// O(log(n))
func binaryGap_Codility(n int) int {
	gap, open1Bit, zeros := 0, false, 0
	if n % 2 != 0 {
		open1Bit = true
	}

	for n != 0 {
		n = n >> 1
		if n & 1 != 0 {
			if !open1Bit {
				open1Bit = true
				zeros = 0
				continue
			}
			if zeros > gap {
				gap = zeros
			}
			zeros = 0
			continue
		}

		zeros++
	}

	return gap
}

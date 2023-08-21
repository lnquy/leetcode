package main

import "fmt"

func main() {
	fmt.Println(isPowerOfThree(27))
	fmt.Println(isPowerOfThree(0))
	fmt.Println(isPowerOfThree(9))
	fmt.Println(isPowerOfThree(45))
}

func isPowerOfThree(n int) bool {
	if n < 1 {
		return false
	}

	for n%3 == 0 {
		n /= 3
	}
	return n == 1
}

// Maximum value which is the power of 3 and fit in int64 is: 4052555153018976267
func isPowerOfThree_PreCalculated(n int) bool {
	return n > 0 && 4052555153018976267%n == 0
}

package main

import "fmt"

func main() {
	fmt.Println(hasAlternatingBits(5))
	fmt.Println(hasAlternatingBits(7))
	fmt.Println(hasAlternatingBits(11))
	fmt.Println(hasAlternatingBits(10))
}

func hasAlternatingBits(n int) bool {
	prevBit := -1
	for n > 0 {
		lastBit := n & 1
		if lastBit != 0 {
			lastBit = 1
		}
		if lastBit == prevBit {
			return false
		}
		prevBit = lastBit
		n >>= 1
	}

	return true
}

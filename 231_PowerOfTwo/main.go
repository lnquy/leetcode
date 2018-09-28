package main

import "fmt"

func main() {
	fmt.Println(isPowerOfTwo(1))
	fmt.Println(isPowerOfTwo(16))
	fmt.Println(isPowerOfTwo(218))
}

func isPowerOfTwo(n int) bool {
	is1Bit := false
	for n > 0 {
		if n&1 == 1 {
			if is1Bit {
				return false
			}
			is1Bit = true
		}
		n >>= 1
	}
	return is1Bit
}

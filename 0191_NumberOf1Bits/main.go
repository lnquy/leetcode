package main

import "fmt"

func main() {
	fmt.Println(hammingWeight(11))
	fmt.Println(hammingWeight(128))
}

func hammingWeight(n uint32) int {
	bits := 0
	for n > 0 {
		bits++
		n &= n - 1
	}
	return bits
}

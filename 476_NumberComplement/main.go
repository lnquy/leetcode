package main

import "fmt"

func main() {
	fmt.Println(findComplement(5))
	fmt.Println(findComplement(1))
	fmt.Println(findComplement(8))
}

func findComplement(num int) int {
	if num == 0 {
		return 1
	}

	i := 1
	for i <= num {
		i <<= 1
	}
	return (i - 1) ^ num
}

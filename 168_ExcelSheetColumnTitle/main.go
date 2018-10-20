package main

import "fmt"

func main() {
	fmt.Println(convertToTitle(1))
	fmt.Println(convertToTitle(28))
	fmt.Println(convertToTitle(701))
}

func convertToTitle(n int) string {
	out := ""
	for n != 0 {
		n--
		out = (string)(n%26+65) + out
		n /= 26
	}
	return out
}

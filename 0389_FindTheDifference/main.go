package main

import "fmt"

func main() {
	fmt.Println(string(findTheDifference("abcd", "abcde")))
}

func findTheDifference(s string, t string) byte {
	out := t[len(s)]
	for i := 0; i < len(s); i++ {
		out += t[i] - s[i]
	}
	return out
}

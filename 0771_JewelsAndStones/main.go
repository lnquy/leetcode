package main

import "fmt"

func main() {
	fmt.Println(numJewelsInStones("aA", "aAAbbbb"))
	fmt.Println(numJewelsInStones("z", "ZZ"))
}

func numJewelsInStones(j string, s string) int {
	m := make(map[rune]struct{}, len(j))

	for _, r := range j {
		m[r] = struct{}{}
	}

	count := 0
	for _, r := range s {
		if _, ok := m[r]; ok {
			count++
		}
	}
	return count
}

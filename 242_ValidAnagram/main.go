package main

import "fmt"

func main() {
	fmt.Println(isAnagram("anagram", "nagaram"))
	fmt.Println(isAnagram("rat", "car"))
}

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	counter := make(map[int]int)
	for i := range s {
		counter[int(s[i]-'a')]++
		counter[int(t[i]-'a')]--
	}
	for _, v := range counter {
		if v != 0 {
			return false
		}
	}
	return true
}

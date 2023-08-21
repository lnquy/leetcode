package main

import "fmt"

func main() {
	fmt.Println(romanToInt("III"))
	fmt.Println(romanToInt("IV"))
	fmt.Println(romanToInt("IX"))
	fmt.Println(romanToInt("LVIII"))
	fmt.Println(romanToInt("MCMXCIV"))
}

func romanToInt(s string) int {
	out, prev := 0, 0
	romans := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	for _, r := range s {
		if prev < romans[r] {
			out = out - (prev * 2) + romans[r]
		} else {
			out = out + romans[r]
		}
		prev = romans[r]
	}
	return out
}

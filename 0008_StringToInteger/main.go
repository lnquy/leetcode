package main

import "fmt"

func main() {
	fmt.Println(myAtoi("42"))
	fmt.Println(myAtoi("   -42"))
	fmt.Println(myAtoi("4193 with words"))
	fmt.Println(myAtoi("words and 987"))
	fmt.Println(myAtoi("-91283472332"))
}

const (
	maxInt32 = 1<<31 - 1
	minInt32 = -1 << 31
)

func myAtoi(str string) int {
	var out int64
	sign, isParsing := "", false

	for _, r := range str {
		if r == ' ' {
			if isParsing { // "  1", " 1 123"
				break
			}
			continue
		}
		if r == '+' || r == '-' {
			if isParsing || sign != "" { // "--0", "++0", "1-123"
				break
			}
			isParsing = true
			sign = string(r)
			continue
		}
		if !(r >= 48 && r <= 57) {
			break
		}

		isParsing = true
		out *= 10
		out += int64(r - 48) // Convert digit's code point to its actual integer value
		if out >= maxInt32 || -out <= minInt32 {
			break
		}
	}

	if sign == "-" {
		out = -out
		if out < minInt32 {
			return minInt32
		}
	}
	if out > maxInt32 {
		return maxInt32
	}
	return int(out)
}

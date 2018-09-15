package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(reverse(123))
	fmt.Println(reverse(-123) == -321)
	fmt.Println(reverse(120) == 21)
}

// O(n)
func reverse(x int) int {
	if x < math.MinInt32 || x > math.MaxInt32 {
		return 0
	}

	isNegative := false
	if x < 0 {
		isNegative = true
		x = -x
	}
	tmp := make([]int, 0, 10)
	step := 1
	for x/10 != 0 {
		tmp = append(tmp, x%10)
		x /= 10
		step *= 10
	}
	tmp = append(tmp, x)

	var out int64
	fmt.Println(tmp)
	for _, v := range tmp {
		out += int64(v * step)
		if out > math.MaxInt32 {
			return 0 // Overflow
		}
		step /= 10
	}
	if isNegative {
		return int(-out)
	}
	return int(out)
}

package main

func minFlips(a int, b int, c int) (flips int) {
	for a > 0 || b > 0 || c > 0 {
		if c&1 == 0 { // Check last bit of c is 0 or 1
			flips += a&1 + b&1
		} else if a&1 == 0 && b&1 == 0 {
			flips++
		}

		a, b, c = a>>1, b>>1, c>>1
	}
	return flips
}

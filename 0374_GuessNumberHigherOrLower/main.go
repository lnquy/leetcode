package main

func guess(num int) int { return -1 }

func guessNumber(n int) int {
	l, r, m := 1, n, 0
	for l <= r {
		m = (l + r) >> 1 // (l+r)/2
		switch guess(m) {
		case 0:
			return m
		case -1:
			r = m - 1
		case 1:
			l = m + 1
		}
	}
	return m
}

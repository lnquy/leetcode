package main

import "fmt"

func main() {
	fmt.Println(addBinary("11", "1"))
}

func addBinary(s1, s2 string) string {
	i, j, carry, out := len(s1)-1, len(s2)-1, '0', ""
	for i >= 0 || j >= 0 {
		b, b1, b2 := rune(-1), '0', '0'
		if i >= 0 {
			b1 = rune(s1[i])
		}
		if j >= 0 {
			b2 = rune(s2[j])
		}

		b, carry = add2Bits(b1, b2, carry)
		out = string(b) + out

		i--
		j--
	}

	if carry == '1' {
		out = string(carry) + out
	}
	return out
}

func add2Bits(b1, b2, carry rune) (rune, rune) {
	if b1 == '0' && b2 == '0' {
		return carry, '0'
	}
	if (b1 == '0' && b2 == '1') || (b1 == '1' && b2 == '0') {
		if carry == '0' {
			return '1', '0'
		}
		return '0', '1'
	}

	if carry == '0' {
		return '0', '1'
	}
	return '1', '1'
}

// Solution from abcman: https://leetcode.com/problems/add-binary/discuss/154785/Go-beats-100-with-single-memory-allocation
// I see it quite smart so I also put it here.
func addBinary_abcman(a string, b string) string {
	iA := len(a) - 1
	iB := len(b) - 1
	sum := 0
	ret := make([]byte, max(len(a), len(b))+1)
	idx := len(ret) - 1
	for iA >= 0 || iB >= 0 {
		if iA >= 0 {
			sum += int(a[iA] - '0')
		}

		if iB >= 0 {
			sum += int(b[iB] - '0')
		}

		ret[idx] = byte(sum&1) + '0'
		sum = sum >> 1
		idx--
		iA--
		iB--
	}

	if sum > 0 {
		ret[0] = '1'
		return string(ret)
	}

	return string(ret[1:])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

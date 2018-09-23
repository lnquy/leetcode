package main

import "fmt"

func main() {
	fmt.Println(convert("PAYPALISHIRING", 3) == "PAHNAPLSIIGYIR")
	fmt.Println(convert("PAYPALISHIRING", 4) == "PINALSIGYAHRPI")
}

func convert(s string, numRows int) string {
	out, n := "", len(s)

	if numRows < 2 || n <= numRows {
		return s
	}

	si := 0
	for si < n { // First row
		out += string(s[si])
		si += 2*numRows - 2
	}

	for i := 1; i < numRows-1; i++ { // Row #2 -> #(numRows-1)
		si = i
		for {
			if si > n-1 {
				break
			}
			out += string(s[si])
			si = (2*numRows - 2 + si) - 2*i
			if si > n-1 {
				break
			}
			out += string(s[si])
			si += 2 * i
		}
	}

	si = numRows - 1
	for si < n { // Last row
		out += string(s[si])
		si += 2*numRows - 2
	}
	return out
}

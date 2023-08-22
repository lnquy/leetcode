package main

import "bytes"

type block struct {
	repeatNo int
	pattern  []byte
}

func decodeString(s string) string {
	var currRepeatNo int
	stack := make([]block, 1, len(s))
	stack[0] = block{
		repeatNo: 1,
		pattern:  []byte{},
	}

	for i := 0; i < len(s); i++ {
		r := s[i]
		switch {
		case r == '0':
			currRepeatNo *= 10
		case r >= '1' && r <= '9':
			currRepeatNo = currRepeatNo*10 + int(r-'0')
		case r == '[':
			stack = append(stack, block{
				repeatNo: currRepeatNo,
				pattern:  []byte{},
			})
			currRepeatNo = 0
		case r == ']':
			lastBlock := stack[len(stack)-1]
			stack = stack[:len(stack)-1] // Pop last block
			stack[len(stack)-1].pattern =
				append(stack[len(stack)-1].pattern, bytes.Repeat(lastBlock.pattern, lastBlock.repeatNo)...)
		default:
			stack[len(stack)-1].pattern = append(stack[len(stack)-1].pattern, r)
		}
	}

	return string(stack[0].pattern)
}

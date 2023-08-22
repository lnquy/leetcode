package main

import (
	"bytes"
	"reflect"
	"unsafe"
)

func removeStars(s string) string {
	stack := make([]byte, 0, len(s))

	for _, r := range stringToBytes(s) {
		if r == '*' {
			stack = stack[:len(stack)-1] // Pop the last item
		} else {
			stack = append(stack, r)
		}
	}

	// Try to eliminate if branching, doesn't worth it, result mostly the same
	// for _, r := range s {
	//		stack = append(stack, r)
	//		if r == '*' {
	//			stack = stack[:len(stack)-2] // Pop the last 2 items
	//		}
	//	}

	return bytes.NewBuffer(stack).String()
}

// Helper func to squeeze some micro performance
func stringToBytes(s string) (b []byte) {
	bh := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
	bh.Data = sh.Data
	bh.Len = sh.Len
	bh.Cap = sh.Len
	return b
}

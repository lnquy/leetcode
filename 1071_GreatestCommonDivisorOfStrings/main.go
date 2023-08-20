package main

import (
	"reflect"
	"unsafe"
)

func gcdOfStrings(str1 string, str2 string) string {
	// Convert string to []byte (unsafe) to prevent string copy every time we passing it to isDivisible()
	str1Bytes, str2Bytes := stringToBytes(str1), stringToBytes(str2)

	maxGCDLen := gcd(len(str1Bytes), len(str2Bytes))
	for i := maxGCDLen; i > 0; i-- {
		gcdStrBytes := str1Bytes[:maxGCDLen]
		if isDivisible(str1Bytes, gcdStrBytes) && isDivisible(str2Bytes, gcdStrBytes) {
			return bytesToString(gcdStrBytes)
		}
	}
	return ""
}

func gcd(a, b int) int {
	min := a
	if b < a {
		min = b
	}
	for i := min; i > 1; i-- {
		if a%i == 0 && b%i == 0 {
			return i
		}
	}
	return 1
}

func isDivisible(parent, sub []byte) bool {
	if len(parent)%len(sub) != 0 {
		return false
	}
	for i := 0; i < len(parent); {
		pattern := parent[i : i+len(sub)]
		if !equalBytes(pattern, sub) {
			return false
		}
		i += len(sub)
	}
	return true
}

// Helper functions to aid performance only, has no affect to the solution correctness
// ---------
func stringToBytes(s string) (b []byte) {
	bh := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
	bh.Data = sh.Data
	bh.Len = sh.Len
	bh.Cap = sh.Len
	return b
}

func bytesToString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

func equalBytes(a, b []byte) bool {
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

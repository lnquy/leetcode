package main

import (
	"bytes"
	"fmt"
	"strconv"
)

func main() {
	in := []byte{'a', 'a', 'b', 'b', 'c', 'c', 'c'}
	fmt.Println(in)
	fmt.Println(compress(in))
	fmt.Println(in)
}

func compress(chars []byte) int {
	lastRune := chars[0]
	counter := 1
	compressed := bytes.NewBuffer(make([]byte, 0, len(chars)))

	for i := 1; i < len(chars); i++ {
		if chars[i] == lastRune {
			counter++
			continue
		}

		_ = compressed.WriteByte(lastRune)
		if counter > 1 {
			_, _ = compressed.WriteString(strconv.Itoa(counter))
		}
		lastRune = chars[i]
		counter = 1
	}

	_ = compressed.WriteByte(lastRune)
	if counter > 1 {
		_, _ = compressed.WriteString(strconv.Itoa(counter))
	}

	// BAD PROBLEM!
	// Cannot reflect the changes of chars to the caller of compress() here.
	// for i, b := range compressed.Bytes() {
	// 	chars[i] = b
	// }

	return compressed.Len()
}

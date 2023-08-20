package main

import (
	"bytes"
)

func mergeAlternately(word1 string, word2 string) string {
	minLen := len(word1)
	if len(word2) < len(word1) {
		minLen = len(word2)
	}
	newString := bytes.NewBuffer(make([]byte, 0, len(word1)+len(word2)))

	for i := 0; i < minLen; i++ {
		_, _ = newString.WriteRune(rune(word1[i]))
		_, _ = newString.WriteRune(rune(word2[i]))
	}

	if len(word1) > len(word2) {
		_, _ = newString.WriteString(word1[minLen:])
	} else if len(word2) > len(word1) {
		_, _ = newString.WriteString(word2[minLen:])
	}
	return newString.String()
}

package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println(uniqueMorseRepresentations([]string{"gin", "zen", "gig", "msg"}))
}

var morse = []string{".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..", ".---", "-.-", ".-..", "--", "-.", "---", ".--.", "--.-", ".-.", "...", "-", "..-", "...-", ".--", "-..-", "-.--", "--.."}

func uniqueMorseRepresentations(words []string) int {
	m := make(map[string]struct{})

	for _, w := range words {
		m[stringToMorse(w)] = struct{}{}
	}
	return len(m)
}

func stringToMorse(s string) string {
	b := bytes.Buffer{}
	for _, r := range s {
		b.WriteString(morse[r-'a'])
	}
	return b.String()
}

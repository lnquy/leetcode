package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	fmt.Println(toGoatLatin("I speak Goat Latin"))
	fmt.Println(toGoatLatin("The quick brown fox jumped over the lazy dog"))
}

func toGoatLatin(s string) string {
	out, suffix, nuWord := bytes.Buffer{}, "", 0

	for _, r := range s {
		if suffix == "" {
			if isVowel(r) {
				suffix = "ma"
				out.WriteRune(r)
			} else {
				suffix = string(r) + "ma"
			}
			continue
		}

		if r == ' ' {
			nuWord++
			out.WriteString(suffix + strings.Repeat("a", nuWord) + " ")
			suffix = ""
		} else {
			out.WriteRune(r)
		}
	}
	nuWord++
	out.WriteString(suffix + strings.Repeat("a", nuWord))
	return out.String()
}

func isVowel(r rune) bool {
	return r == 'a' || r == 'e' || r == 'u' || r == 'i' || r == 'o' ||
		r == 'A' || r == 'E' || r == 'U' || r == 'I' || r == 'O'
}

package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(mostCommonWord("Bob hit a ball, the hit BALL flew far after it was hit.", []string{"hit"}))
}

func mostCommonWord(paragraph string, banned []string) string {
	bannedWords := make(map[string]struct{}, len(banned))
	for _, v := range banned {
		bannedWords[v] = struct{}{}
	}

	words := strings.FieldsFunc(paragraph, func(r rune) bool {
		return r == ' ' || r == '!' || r == '?' || r == '\'' || r == ',' || r == ';' || r == '.'
	})

	wordList := make(map[string]int)
	out, maxExists := "", 0
	for _, w := range words {
		w = strings.ToLower(w)
		if _, ok := bannedWords[w]; ok {
			continue
		}
		wordList[w]++
		if v := wordList[w]; v > maxExists {
			maxExists = v
			out = w
		}
	}
	return out
}

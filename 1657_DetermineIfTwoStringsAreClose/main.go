package main

import "sort"

func closeStrings(word1 string, word2 string) bool {
	counterFunc := func(word string) (runes, occurrences [26]int) {
		for i := range word {
			runes[word[i]-'a'] = 1
			occurrences[word[i]-'a'] += 1
		}
		sort.Ints(occurrences[:])
		return runes, occurrences
	}

	runes1, occurrences1 := counterFunc(word1)
	runes2, occurrences2 := counterFunc(word2)
	return runes1 == runes2 && occurrences1 == occurrences2
}

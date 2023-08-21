package main

func maxVowels(s string, k int) (max int) {
	slidingVowels := 0
	for i := 0; i < k; i++ {
		if isVowel(s[i]) {
			slidingVowels++
		}
	}
	if len(s) == k {
		return slidingVowels
	}

	max = slidingVowels

	for i := k; i < len(s); i++ {
		if isVowel(s[i-k]) {
			slidingVowels--
		}
		if isVowel(s[i]) {
			slidingVowels++
		}
		if slidingVowels > max {
			max = slidingVowels
		}
	}

	return max
}

func isVowel(s byte) bool {
	return s == 'a' || s == 'e' || s == 'i' || s == 'o' || s == 'u'
}

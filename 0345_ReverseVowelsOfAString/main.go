package main

func reverseVowels(s string) string {
	sb := []byte(s)
	rightIdx := len(sb) - 1
	for leftIdx := 0; leftIdx <= rightIdx; leftIdx++ {
		if !isVowel(sb[leftIdx]) {
			continue
		}
		rightIdx = getVowelRightIdx(sb, leftIdx, rightIdx)
		if leftIdx >= rightIdx { // End early if we have less vowels on the right side
			return string(sb)
		}
		sb[leftIdx], sb[rightIdx] = sb[rightIdx], sb[leftIdx]
		rightIdx--
	}
	return string(sb)
}

func isVowel(b byte) bool {
	return b == 'a' || b == 'e' || b == 'i' || b == 'o' || b == 'u' ||
		b == 'A' || b == 'E' || b == 'I' || b == 'O' || b == 'U'
}

func getVowelRightIdx(sb []byte, boundIdx, rightIdx int) int {
	for ; rightIdx >= boundIdx; rightIdx-- {
		if isVowel(sb[rightIdx]) {
			return rightIdx
		}
	}
	return -1
}

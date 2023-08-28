package main

import (
	"math"
	"sort"
)

func successfulPairs(spells []int, potions []int, success int64) (result []int) {
	sort.Ints(potions)
	result = make([]int, len(spells), len(spells))

	for i, spell := range spells {
		minPotion := int(math.Ceil(float64(success) / float64(spell)))
		result[i] = countBiggerThan(potions, minPotion)
	}
	return result
}

func countBiggerThan(s []int, min int) int {
	if s[len(s)-1] < min {
		return 0
	}

	left, right := 0, len(s)
	for left <= right {
		mid := (left + right) / 2
		if s[mid] == min {
			if mid == 0 || (mid > 0 && s[mid] != s[mid-1]) {
				return len(s) - mid
			}
			right = mid - 1
			continue
		}

		if s[mid] < min {
			left = mid + 1
			continue
		}

		if s[mid] > min {
			if mid == 0 || (mid > 0 && s[mid-1] < min) {
				return len(s) - mid
			}
			right = mid - 1
			continue
		}
	}
	return 0
}

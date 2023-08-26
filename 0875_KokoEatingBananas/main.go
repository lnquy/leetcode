package main

import (
	"math"
)

func minEatingSpeed(piles []int, h int) (eatingRate int) {
	min, max := 1, biggestPile(piles)
	var eatingHours int

	for min < max {
		eatingRate = (min + max) / 2
		eatingHours = 0

		for _, p := range piles {
			eatingHours += int(math.Ceil(float64(p) / float64(eatingRate)))
			if eatingHours > h { // Eating too slow
				min = eatingRate + 1
				break
			}
		}

		if eatingHours <= h { // Eating too fast
			max = eatingRate
		}
	}
	return max
}

func biggestPile(piles []int) (max int) {
	for _, p := range piles {
		if p > max {
			max = p
		}
	}
	return max
}

package main

func largestAltitude(gain []int) (highestAlt int) {
	var lastAlt int

	for _, n := range gain {
		lastAlt += n
		if lastAlt > highestAlt {
			highestAlt = lastAlt
		}
	}
	return highestAlt
}

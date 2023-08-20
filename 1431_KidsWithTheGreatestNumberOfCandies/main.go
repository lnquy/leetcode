package main

func kidsWithCandies(candies []int, extraCandies int) []bool {
	max := candies[0]
	for i := 1; i < len(candies); i++ {
		if max < candies[i] {
			max = candies[i]
		}
	}
	result := make([]bool, len(candies), len(candies))
	for i := 0; i < len(candies); i++ {
		if candies[i]+extraCandies >= max {
			result[i] = true
		}
	}
	return result
}

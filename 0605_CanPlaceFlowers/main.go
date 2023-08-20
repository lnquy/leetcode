package main

func canPlaceFlowers(flowerbed []int, n int) bool {
	flLen := len(flowerbed)
	switch flLen {
	case 1:
		if flowerbed[0] == 0 {
			n--
		}
	case 2:
		if flowerbed[0] == 0 && flowerbed[1] == 0 {
			n--
		}
	default:
		// First item in array. In case [0, 0, ...], always prioritize placing at the first position
		if flowerbed[0] == 0 && flowerbed[1] == 0 {
			n--
			flowerbed[0] = 1
		}

		// All items in between the array
		for i := 1; i < flLen-1; i++ {
			if flowerbed[i] == 1 {
				continue
			}
			if flowerbed[i-1] == 0 && flowerbed[i+1] == 0 {
				n--
				flowerbed[i] = 1
			}
		}

		// Last item in array. In case [..., 0, 0], always prioritize placing at the last position
		if flowerbed[flLen-1] == 0 && flowerbed[flLen-2] == 0 {
			n--
			// flowerbed[flLen-1] = 1 // Comment to save a little time
		}
	}

	return n <= 0
}

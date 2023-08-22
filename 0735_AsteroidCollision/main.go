package main

func asteroidCollision(asteroids []int) []int {
	stack := make([]int, 0, len(asteroids))

	for _, currAst := range asteroids {
		if len(stack) == 0 {
			stack = append(stack, currAst)
			continue
		}

		if currAst > 0 { // Current asteroid is positive.
			// Positive asteroid move to the right and doesn't care about the last asteroid in the stack
			stack = append(stack, currAst)
		} else { // Current asteroid is negative
			lastAst := stack[len(stack)-1]
			if lastAst < 0 { // No collide
				stack = append(stack, currAst)
				continue
			}
			// Collide with positive asteroids in stack until there's non left
			shouldAddCurrAstToStack := false
			for lastAst > 0 {
				cabs, labs := abs(int32(currAst)), abs(int32(lastAst))
				if cabs < labs { // currAst got destroyed
					shouldAddCurrAstToStack = false
					break
				} else if cabs > labs { // lastAst destroyed
					shouldAddCurrAstToStack = true
					stack = stack[:len(stack)-1]
					if len(stack) <= 0 {
						break
					}
					lastAst = stack[len(stack)-1]
				} else if cabs == labs { // both asts destroyed
					stack = stack[:len(stack)-1] // lastAst destroyed
					shouldAddCurrAstToStack = false
					break
				}
			}
			if shouldAddCurrAstToStack {
				stack = append(stack, currAst)
			}
		}
	}

	return stack
}

func abs(x int32) int {
	y := x >> 31
	return int((x ^ y) - y)
}

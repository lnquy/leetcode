package main

import "fmt"

func main() {
	fmt.Println(lemonadeChange([]int{5, 5, 5, 10, 20}))
	fmt.Println(lemonadeChange([]int{5, 5, 10}))
	fmt.Println(lemonadeChange([]int{10, 10}))
	fmt.Println(lemonadeChange([]int{5, 5, 10, 10, 20}))
}

func lemonadeChange(bills []int) bool {
	fives, tens := 0, 0

	for i := 0; i < len(bills); i++ {
		if bills[i] == 5 {
			fives++
			continue
		}
		if bills[i] == 10 {
			if fives == 0 {
				return false
			}
			fives--
			tens++
			continue
		}
		if bills[i] == 20 {
			if tens == 0 {
				if fives < 3 {
					return false
				}
				fives -= 3
				continue
			}
			if fives == 0 {
				return false
			}
			tens--
			fives--
		}
	}
	return true
}

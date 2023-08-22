package main

type RecentCounter struct {
	arr []int
}

func Constructor() RecentCounter {
	return RecentCounter{[]int{}}
}

func (c *RecentCounter) Ping(t int) int {
	c.arr = append(c.arr, t)

	counter := 0
	for i := len(c.arr) - 1; i >= 0; i-- {
		if c.arr[i] >= t-3000 {
			counter++
		} else {
			break
		}
	}
	return counter
}

// Faster solution but not as intuitive
// type RecentCounter struct {
// 	k   int
// 	arr []int
// }
//
// func Constructor() RecentCounter {
// 	return RecentCounter{0, []int{}}
// }
//
// func (c *RecentCounter) Ping(t int) int {
// 	c.arr = append(c.arr, t)
//
// 	for c.arr[c.k] < t-3000 {
// 		c.k++
// 	}
// 	return len(c.arr) - c.k
// }

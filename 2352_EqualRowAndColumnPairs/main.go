package main

import (
	"bytes"
	"strconv"
)

func equalPairs(grid [][]int) (counter int) {
	n := len(grid)
	mRows, mColumns := make(map[string]int, n), make(map[string]int, n)
	rowStr := bytes.NewBuffer(make([]byte, 0, n*3))
	clmStr := bytes.NewBuffer(make([]byte, 0, n*3))

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			_, _ = rowStr.WriteString(strconv.Itoa(grid[i][j]))
			_ = rowStr.WriteByte(':')
			_, _ = clmStr.WriteString(strconv.Itoa(grid[j][i]))
			_ = clmStr.WriteByte(':')
		}
		mRows[rowStr.String()]++
		mColumns[clmStr.String()]++
		rowStr.Reset()
		clmStr.Reset()
	}

	for k, v := range mRows {
		if occ, ok := mColumns[k]; ok {
			counter += v * occ
		}
	}
	return counter
}

// Unoptimized version
// func equalPairs(grid [][]int) int {
// 	n := len(grid)
// 	mRows, mColumns := make(map[string]int, n), make(map[string]int, n)
//
// 	for i := 0; i < n; i++ {
// 		rowStr := bytes.NewBuffer(make([]byte, 0, n*5))
// 		clmStr := bytes.NewBuffer(make([]byte, 0, n*5))
// 		for j := 0; j < n; j++ {
// 			_, _ = rowStr.WriteString(fmt.Sprintf("%d:", grid[i][j]))
// 			_, _ = clmStr.WriteString(fmt.Sprintf("%d:", grid[j][i]))
// 		}
// 		mRows[rowStr.String()]++
// 		mColumns[clmStr.String()]++
// 	}
//
// 	var counter int
// 	for k, v := range mRows {
// 		if occ, ok := mColumns[k]; ok {
// 			counter += v * occ
// 		}
// 	}
//
// 	return counter
// }

package problem

import (
	"fmt"
	"math/rand"
)

func isBadVersion(version int) bool {
	return rand.Intn(version)%2 == 1
}

func print2DGrid(grid [][]int) {
	m, n := len(grid), len(grid[0])

	for i := 0; i < m; i++ {
		if grid[i] == nil {
			fmt.Println()
			continue
		}

		for j := 0; j < n; j++ {
			fmt.Print(grid[i][j], "\t")
		}
		fmt.Println()
	}
	fmt.Println()
}

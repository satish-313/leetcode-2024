package april

import (
	"fmt"
)

func islandPerimeter(grid [][]int) int {
	peri := 0
	row := len(grid)
	col := len(grid[0])
	dir := [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if grid[i][j] == 1 {
				for _, v := range dir {
					ri := v[0] + i
					cj := v[1] + j

					if ri < 0 || ri >= row || cj < 0 || cj >= col || grid[ri][cj] == 0 {
						peri += 1
					}
				}
			}
		}
	}

	return peri
}

func Day18() {
	q := [][][]int{
		{[]int{0, 1, 0, 0}, []int{1, 1, 1, 0}, []int{0, 1, 0, 0}, []int{1, 1, 0, 0}},
		{[]int{1}},
		{[]int{1, 0}},
	}

	for _, v := range q {
		fmt.Println(islandPerimeter(v))
	}
}

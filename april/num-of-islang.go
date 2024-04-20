package april

import "fmt"

func numIslands(grid [][]byte) int {
	ti := 0

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '1' {
				ti += 1
				dfs19(&grid, i, j)
			}
		}
	}

	return ti
}

func dfs19(grid *[][]byte, r, c int) {
	(*grid)[r][c] = '0'
	dir := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	row := len(*grid)
	col := len((*grid)[0])

	for i := 0; i < len(dir); i++ {
		ri := r + dir[i][0]
		cj := c + dir[i][1]

		if ri < 0 || ri >= row || cj < 0 || cj >= col || (*grid)[ri][cj] == '0' {
			continue
		}
		dfs19(grid, ri, cj)
	}
}

func Day19() {
	grid := [][]byte{
		{'1', '1', '1', '1', '0'},
		{'1', '1', '0', '1', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '0', '0', '0'},
	}
	fmt.Println(numIslands(grid))
}

package april

import "fmt"

func findFarmland(land [][]int) [][]int {
	res := [][]int{}

	for i := 0; i < len(land); i++ {
		for j := 0; j < len(land[0]); j++ {
			if land[i][j] == 1 {
				temp := []int{0, 0, 0, 0}
				temp[0] = i
				temp[1] = j
				temp[2] = i
				temp[3] = j
				dfs20(&land, i, j, &temp[3], &temp[2])
				res = append(res, temp)
			}
		}
	}

	return res
}

func dfs20(l *[][]int, r, c int, cM, rM *int) {
	if r < 0 || c < 0 || r >= len(*l) || c >= len((*l)[0]) || (*l)[r][c] == 0 {
		return
	}

	(*l)[r][c] = 0
	if c >= *cM && r >= *rM {
		*cM = c
		*rM = r
	}
	dfs20(l, r+1, c, cM, rM)
	dfs20(l, r, c+1, cM, rM)
}

func Day20() {
	land := [][]int{{1, 0, 0}, {0, 1, 1}, {0, 1, 1}}
	fmt.Println(findFarmland(land))
}

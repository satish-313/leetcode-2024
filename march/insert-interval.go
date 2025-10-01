package march

import (
	"fmt"
)

type day17 struct {
	arr [][]int
	ins []int
}

func insert(in [][]int, ni []int) [][]int {
	res := [][]int{}
	n := len(in)

	if n == 0 {
		res = append(res, ni)
		return res
	}

	if in[n-1][1] < ni[0] {
		in = append(in, ni)
		return in
	}

	if in[0][0] > ni[1] {
		res = append(res, ni)
		for i := 0; i < n; i++ {
			res = append(res, in[i])
		}
		return res
	}

	for i := 0; i < n; i++ {
		v := in[i]
		if v[1] < ni[0] {
			res = append(res, v)
		} else if v[0] > ni[1] {
			if in[i-1][1] < ni[0] {
				res = append(res, ni)
			}
			res = append(res, v)
		} else if v[1] >= ni[0] {
			t := make([]int, 2)
			if v[0] < ni[0] {
				t[0] = v[0]
			} else {
				t[0] = ni[0]
			}

			for ; i < n; i++ {
				v2 := in[i]

				if ni[1] < v2[0] {
					t[1] = ni[1]
					i -= 1
					break
				} else if ni[1] <= v2[1] {
					t[1] = v2[1]
					break
				}

				if i == n-1 && ni[1] > v2[1] {
					t[1] = ni[1]
				}
			}
			res = append(res, t)
		}
	}

	return res
}

func Day17() {
	q := []day17{
		{arr: [][]int{{3, 5}, {12, 15}}, ins: []int{6, 6}},
		// {arr: [][]int{{1, 3}, {6, 9}}, ins: []int{2, 5}},
		// {arr: [][]int{{1, 5}}, ins: []int{2, 7}},
		// {arr: [][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}}, ins: []int{4, 8}},
	}

	for _, v := range q {
		fmt.Println(insert(v.arr, v.ins))
	}
}

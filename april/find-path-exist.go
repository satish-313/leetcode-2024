package april

import (
	"fmt"
)

type day21 struct {
	n           int
	edges       [][]int
	source      int
	destination int
}

func validPath(n int, edges [][]int, source int, destination int) bool {
	if n == 1 {
		return true
	}
	visit := make([]bool, n)
	m := make(map[int][]int)

	for _, v := range edges {
		m[v[0]] = append(m[v[0]], v[1])
		m[v[1]] = append(m[v[1]], v[0])
	}

	stack := m[source]
	for len(stack) > 0 {
		length := len(stack)
		for _, v := range stack {
			if v == destination {
				return true
			}
			if visit[v] {
				continue
			}
			visit[v] = true
			next := m[v]
			for _, val := range next {
				if !visit[val] {
					stack = append(stack, val)
				}
			}
		}
		stack = stack[length:]
	}

	return false
}

func Day21() {
	q := []day21{
		{n: 3, edges: [][]int{{0, 1}, {1, 2}, {2, 0}}, source: 0, destination: 2},
		{n: 6, edges: [][]int{{0, 1}, {0, 2}, {3, 5}, {5, 4}, {4, 3}}, source: 0, destination: 5},
		{n: 10, edges: [][]int{{4, 3}, {1, 4}, {4, 8}, {1, 7}, {6, 4}, {4, 2}, {7, 4}, {4, 0}, {0, 9}, {5, 4}}, source: 5, destination: 9},
	}
	for _, v := range q {
		fmt.Println(validPath(v.n, v.edges, v.source, v.destination))
	}
}

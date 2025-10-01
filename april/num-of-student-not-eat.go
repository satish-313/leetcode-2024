package april

import (
	"fmt"
)

type day8 struct {
	std  []int
	sdwh []int
}

func countStudents(students []int, sandwiches []int) int {
	arr := [2]int{0, 0}

	for _, v := range students {
		arr[v]++
	}

	for len(students) > 0 {
		if students[0] == sandwiches[0] {
			arr[students[0]]--
			students = students[1:]
			sandwiches = sandwiches[1:]
		} else {
			if arr[students[0]] == 0 || arr[sandwiches[0]] == 0 {
				break
			}
			students = append(students, students[0])
			students = students[1:]
		}
	}
	return len(students)
}

func Day8() {
	q := []day8{
		{std: []int{1, 1, 0, 0}, sdwh: []int{0, 1, 0, 1}},
		{std: []int{1, 1, 1, 0, 0, 1}, sdwh: []int{1, 0, 0, 0, 1, 1}},
	}

	for _, v := range q {
		fmt.Println(countStudents(v.std, v.sdwh))
	}
}

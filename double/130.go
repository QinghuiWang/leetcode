package double

import "sort"

// https://leetcode.cn/contest/biweekly-contest-130/

//1
func satisfiesConditions(grid [][]int) bool {

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if i+1 < len(grid) && grid[i][j] != grid[i+1][j] {
				return false
			}
			if j+1 < len(grid[i]) && grid[i][j] == grid[i][j+1] {
				return false
			}
		}
	}
	return true
}

//2
func maxPointsInsideSquare(points [][]int, s string) int {

	p := make([]P, 0, len(points))
	for index, v := range points {
		max := Max(Abs(v[0]), Abs(v[1]))
		p = append(p, P{index: index, max: max})
	}
	sort.Slice(p, func(i, j int) bool {
		return p[i].max < p[j].max
	})

	result := 0
	l := 0
	totalTag := map[byte]bool{}
	tag := map[byte]bool{}
	for _, point := range p {

		if point.max > l {
			result += len(tag)
			clear(tag)
		}
		if totalTag[s[point.index]] {
			return result
		} else {
			totalTag[s[point.index]] = true
		}
		l = point.max
		if tag[s[point.index]] {
			return result
		} else {
			tag[s[point.index]] = true
		}
	}
	result += len(tag)
	return result
}

type P struct {
	index int
	max   int
}

func Abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func Max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

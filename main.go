package main

import (
	"slices"
	"sort"
	"strings"
)

func main() {
	reverseWords("a   example good")
}

func reverseWords(s string) string {
	list := strings.Split(s, " ")

	resBuilder := strings.Builder{}

	for i := 0; i < len(list); i++ {
		v := list[len(list)-1-i]
		if v == " " || v == "" {
			continue
		} else {
			resBuilder.WriteString(strings.TrimSpace(v))
			resBuilder.WriteString(" ")
		}
	}
	return strings.TrimSpace(resBuilder.String())
}

func satisfiesConditions(grid [][]int) bool {

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if i+1 < len(grid) && grid[i][j] == grid[i+1][j] {
				return false
			}
			if j+1 < len(grid[i]) && grid[i][j] != grid[i][j+1] {
				return false
			}
		}
	}
	return true
}

type P struct {
	index int
	max   int
}

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
		if totalTag[s[point.index]] {
			return result
		} else {
			totalTag[s[point.index]] = true
		}

		if point.max > l {
			result += len(tag)
			clear(tag)
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


func Max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func findPermutationDifference(s string, t string) int {
	sum:=0
	for index, v := range s {
		idx:=strings.IndexRune(t,v)
		sum+=Abs(index-idx)
	}
	return sum
}

func Abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}


// 
func maxScore(grid [][]int) int {
	dp:=make([][]int,len(grid))
	for i := 0; i < len(dp); i++ {
		dp[i]=make([]int,len(grid[0]))
	}
	res:=-1000000000
    dp[0][0]=0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
            if i==0 && j==0 {
                continue
            }
			ans:=-10000000
			for m := 0; m < j; m++ {
                if(dp[i][m]>0){
                    ans=Max(ans,grid[i][j]-grid[i][m]+dp[i][m])
                }else{
                    ans=Max(ans,grid[i][j]-grid[i][m])
                }
				
			}
			for m := 0; m < i; m++ {
                if(dp[m][j]>0){
                   ans=Max(ans,grid[i][j]-grid[m][j]+dp[m][j])
                }else{
                    ans=Max(ans,grid[i][j]-grid[m][j])a
                }
				
			}
			dp[i][j]=ans
			res=Max(res,ans)
		}
	}
	return res
}
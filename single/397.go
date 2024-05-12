package single

// https://leetcode.cn/contest/weekly-contest-397
// 1
func findPermutationDifference(s string, t string) int {
	sum:=0
	for index, v := range s {
		idx:=strings.IndexRune(t,v)
		sum+=Abs(index-idx)
	}
	return sum
}

// 2
func maximumEnergy(energy []int, k int) int {
    dp:=make([]int,len(energy))
	for index,v := range energy {
		if index-k>=0 {
			dp[index]=Max(dp[index-k]+energy[index],energy[index])
		}else{
			dp[index]=v
		}
	}

    max:=-1001
    for i:=len(dp)-k;i<len(dp);i++ {
        max=Max(max,dp[i])
    }
	return max
}


// 3
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
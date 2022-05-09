package dp

/*
官方的定义是指把多阶段过程转化为一系列单阶段问题，利用各阶段之间的关系，逐个求解。概念中的各阶段之间的关系，其实指的就是状态转移方程
一般看到的状态转移方程，基本这样:
opt ：指代特殊的计算逻辑，通常为 max or min。
i,j,k 都是在定义DP方程中用到的参数。
dp[i] = opt(dp[i-1])+1
dp[i][j] = w(i,j,k) + opt(dp[i-1][k])
dp[i][j] = opt(dp[i-1][j] + xi, dp[i][j-1] + yj, ...)

经典爬楼梯问题：
一次爬1或2个台阶，爬到 n 阶有多少种方法
解题思路：
1.爬到1阶有1种，爬到2阶有两种 1 2
2.爬到三阶有三种，1 1 1/ 1 2 / 2 1
3.爬到n阶有 dp(n-1) + dp(n-2) 中，n阶也就是n-1阶爬1个或n-2爬2个
*/

func climbStairs(n int) int {
	if n == 1 {
		return 1
	}
	dp := make([]int, n+1)
	dp[1] = 1
	dp[2] = 2
	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}

	return dp[n]
}

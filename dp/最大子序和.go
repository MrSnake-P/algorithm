package dp

import "algorithm/util"

/*
最大子序和
给定一个整数数组 nums ，找到一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。
输入: [-2,1,-3,4,-1,2,1,-5,4],
输出: 6
解释: 连续子数组 [4,-1,2,1] 的和最大，为 6。
*/

/*
解题思路：
1.要想得到dp[i]一定和num[i]有关
2.dp[i]和dp[i-1]很可能只差一个num[i] => dp[i] = dp[i-1]+num[i]
3.但是如果dp[i-1]为负的，num[i]才是最大的 => dp[i] = max(num[i], dp[i-1]+num[i])
4.注意dp[i]并不是要求的最终结果，max(dp[0], dp[1], ...dp[i])
*/

func maxSubArray(nums []int) int {
	if len(nums) < 1 {
		return 0
	}
	// 首先需要根据已有的已有的状态进行推导
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		dp[i] = util.Max(nums[i], dp[i-1]+nums[i])
	}
	r := -1 << 31
	for _, v := range dp {
		r = util.Max(v, r)
	}
	return r
}

/*
优化后
*/

func maxSubArray2(nums []int) int {
	if len(nums) < 1 {
		return 0
	}
	r := -1 << 31
	// 首先需要根据已有的已有的状态进行推导
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		dp[i] = util.Max(nums[i], dp[i-1]+nums[i])
		r = util.Max(dp[i], r)
	}

	return r
}

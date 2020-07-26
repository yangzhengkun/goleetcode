package _209_minimum_size_subarray_sum

import "math"

//给定一个含有n个正整数的数组和一个正整数s,
//找出该数组中满足其和 ≥ s 的长度最小的连续子数组,并返回其长度.
//如果不存在符合条件的子数组,返回0
//示例:
//输入:s = 7, nums = [2,3,1,2,4,3]
//输出:2
//解释:子数组[4,3]是该条件下的长度最小的子数组。
//

//滑动窗口
func minSubArrayLen(s int, nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	min := func(x, y int) int {
		if x < y {
			return x
		}
		return y
	}
	ans := math.MaxInt32
	left, right := 0, 0 // nums[left...right]区间的元素之和>=s
	sum := 0
	for right < n {
		sum += nums[right]
		for sum >= s {
			sum -= nums[left]
			ans = min(ans, right-left+1)
			left++
		}
		right++
	}
	if ans == math.MaxInt32 {
		return 0
	}
	return ans
}

package _239_sliding_window_maximum

//给定一个数组 nums，有一个大小为 k 的滑动窗口从数组的最左侧移动到数组的最右侧。你只可以看到在滑动窗口内的 k 个数字。滑动窗口每次只向右移动一位。
//
//返回滑动窗口中的最大值。
//
//进阶：
//
//你能在线性时间复杂度内解决此题吗？
//
//示例:
//
//输入: nums = [1,3,-1,-3,5,3,6,7], 和 k = 3
//输出: [3,3,5,5,6,7]
//解释:
//
//  滑动窗口的位置                最大值
//---------------               -----
//[1  3  -1] -3  5  3  6  7       3
// 1 [3  -1  -3] 5  3  6  7       3
// 1  3 [-1  -3  5] 3  6  7       5
// 1  3  -1 [-3  5  3] 6  7       5
// 1  3  -1  -3 [5  3  6] 7       6
// 1  3  -1  -3  5 [3  6  7]      7
//

func maxSlidingWindow(nums []int, k int) []int {
	if nums == nil || k <= 0 {
		return []int{}
	}
	var window []int
	var res []int
	for i, val := range nums {
		if i >= k && window[0] <= i-k {
			window = window[1:]
		}
		for len(window) > 0 && nums[window[len(window)-1]] <= val {
			window = window[:len(window)-1]
		}
		window = append(window, i)
		if i >= k-1 {
			res = append(res, nums[window[0]])
		}
	}
	return res
}

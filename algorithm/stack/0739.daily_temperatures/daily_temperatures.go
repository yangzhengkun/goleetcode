package _739_daily_temperatures

//请根据每日 气温 列表，重新生成一个列表。对应位置的输出为：要想观测到更高的气温，至少需要等待的天数。如果气温在这之后都不会升高，请在该位置用 0 来代替。
//
//例如，给定一个列表 temperatures = [73, 74, 75, 71, 69, 72, 76, 73]，你的输出应该是 [1, 1, 4, 2, 1, 1, 0, 0]。
//
//提示：气温 列表长度的范围是 [1, 30000]。每个气温的值的均为华氏度，都是在 [30, 100] 范围内的整数。
//

// 给定一个数组，找到每个位置i右边第一个比arr[i]大的位置

// 单调栈结构:从栈底到栈顶是递减的.每一次弹出栈顶位置i的时候,表示此时遍历到的元素位置j,就是第一个比位置i元素大的数
func dailyTemperatures(T []int) []int {
	var res = make([]int, len(T))
	var stack []int
	for i := 0; i < len(T); i++ {
		for len(stack) > 0 && T[i] > T[stack[len(stack)-1]] {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			res[top] = i - top
		}
		stack = append(stack, i)
	}
	return res
}

package _167_two_sum_ii

//给定一个已按照升序排列 的有序数组，找到两个数使得它们相加之和等于目标数。
//
//函数应该返回这两个下标值 index1 和 index2，其中 index1 必须小于 index2。
//
//说明:
//
//返回的下标值（index1 和 index2）不是从零开始的。
//你可以假设每个输入只对应唯一的答案，而且你不可以重复使用相同的元素。
//示例:
//
//输入: numbers = [2, 7, 11, 15], target = 9
//输出: [1,2]
//解释: 2 与 7 之和等于目标数 9 。因此 index1 = 1, index2 = 2 。
//

//对撞指针
func twoSum(numbers []int, target int) []int {
	left, right := 0, len(numbers)-1
	var sum int
	for left < right {
		//一定要清晰l和r的意义,当l==r时已经是无效的区间
		sum = numbers[left] + numbers[right]
		if sum == target {
			return []int{left + 1, right + 1}
		} else if sum < target {
			left++
		} else {
			right--
		}
	}
	return []int{}
}

//hash
func twoSum1(numbers []int, target int) []int {
	var result = make(map[int]int)
	var ret []int
	for i := 0; i < len(numbers); i++ {
		if index, ok := result[target-numbers[i]]; ok {
			ret = []int{index + 1, i + 1}
		}
		result[numbers[i]] = i
	}
	return ret
}

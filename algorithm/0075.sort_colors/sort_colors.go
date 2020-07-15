package _075_sort_colors

//给定一个包含红色、白色和蓝色，一共 n 个元素的数组，原地对它们进行排序，使得相同颜色的元素相邻，并按照红色、白色、蓝色顺序排列。
//
//此题中，我们使用整数 0、 1 和 2 分别表示红色、白色和蓝色。
//
//注意:
//不能使用代码库中的排序函数来解决这道题。
//
//示例:
//
//输入: [2,0,2,1,1,0]
//输出: [0,0,1,1,2,2]
//

//三路快排
func sortColors(nums []int) {
	if nums == nil {
		return
	}
	var zero = -1       // nums[0...zeros] == 0
	var two = len(nums) // nums[two...n-1] == 2
	for i := 0; i < two; {
		if nums[i] == 0 {
			zero++
			nums[zero], nums[i] = nums[i], nums[zero]
			i++
		} else if nums[i] == 1 {
			i++
		} else {
			// nums[i] == 2
			two--
			nums[i], nums[two] = nums[two], nums[i]
		}
	}
}

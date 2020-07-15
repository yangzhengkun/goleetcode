package _283_move_zeroes

//双指针:快慢指针
func moveZeroes(nums []int) {
	var index = 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[index] = nums[i]
			index++
		}
	}
	//index之后的位置上要置为0
	for i := index; i < len(nums); i++ {
		nums[i] = 0
	}
}

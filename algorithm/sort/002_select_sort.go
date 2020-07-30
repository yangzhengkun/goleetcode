package sort

func SelectSort(arr []int) []int {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		min := i // 每次循环从[i+1,n-1]之间选出最小值,并与arr[i]交换
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[min] {
				min = j
			}
		}
		arr[i], arr[min] = arr[min], arr[i]
	}
	return arr
}

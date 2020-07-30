package sort

func InsertSort(arr []int) []int {
	n := len(arr)
	for i := 0; i < n; i++ {
		j := i - 1
		cur := arr[i]
		for j > 0 && arr[j] > cur {
			arr[j+1] = arr[j] // j位置元素往后移动一个
			j -= 1
		}
		arr[j] = cur
	}
	return arr
}

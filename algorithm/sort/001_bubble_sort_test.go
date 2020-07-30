package sort

import (
	"fmt"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	arr := []int{5}
	fmt.Println(BubbleSort(arr))

	arr = []int{1, 2, 3, 5, 2, 6, 13, 3}
	fmt.Println(BubbleSort(arr))
}

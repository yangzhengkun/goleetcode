package _075_sort_colors

import (
	"fmt"
	"testing"
)

// go test -v .
func TestSortColors(t *testing.T) {
	//var nums = []int{2, 0, 2, 1, 1, 0}
	var nums = []int{2, 0, 2, 1, 1, 0}
	sortColors(nums)
	fmt.Println(nums)
}

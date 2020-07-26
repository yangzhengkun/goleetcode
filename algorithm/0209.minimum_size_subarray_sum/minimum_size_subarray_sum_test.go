package _209_minimum_size_subarray_sum

import (
	"fmt"
	"testing"
)

func TestMinimumSizeSubarray(t *testing.T) {
	nums := []int{2, 3, 1, 2, 4, 3}
	s := 7
	fmt.Println(minSubArrayLen(s, nums))
}

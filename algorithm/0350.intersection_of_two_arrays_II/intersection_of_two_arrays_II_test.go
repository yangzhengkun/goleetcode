package _350_intersection_of_two_arrays_II

import (
	"fmt"
	"testing"
)

func TestIntersect(t *testing.T) {
	nums1 := []int{4, 9, 5}
	nums2 := []int{9, 4, 9, 8, 4}
	fmt.Println(intersect(nums1, nums2))
}

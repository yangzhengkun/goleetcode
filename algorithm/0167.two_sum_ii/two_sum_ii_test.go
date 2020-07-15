package _167_two_sum_ii

import (
	"fmt"
	"testing"
)

func TestTwoSumII(t *testing.T) {
	var numbers = []int{2, 7, 11, 15}
	ret := twoSum(numbers, 9)
	fmt.Println(ret)
}

func TestTwoSumII1(t *testing.T) {
	var numbers = []int{2, 7, 11, 15}
	ret := twoSum1(numbers, 9)
	fmt.Println(ret)
}

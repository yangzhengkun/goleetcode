package _084_largest_rectangle_in_histogram

import "fmt"

//给定 n 个非负整数，用来表示柱状图中各个柱子的高度。每个柱子彼此相邻，且宽度为 1 。
//
//求在该柱状图中，能够勾勒出来的矩形的最大面积。给定 n 个非负整数，用来表示柱状图中各个柱子的高度。每个柱子彼此相邻，且宽度为 1 。
//
func largestRectangleArea(heights []int) int {
	var max = func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}

	n := len(heights)
	left, right := make([]int, n), make([]int, n)
	for i := 0; i < n; i++ {
		right[i] = n
	}
	var stack []int
	for i := 0; i < n; i++ {
		for len(stack) > 0 && heights[stack[len(stack)-1]] >= heights[i] {
			right[stack[len(stack)-1]] = i
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			left[i] = -1
		} else {
			left[i] = stack[len(stack)-1]
		}
		stack = append(stack, i)
	}
	ans := 0
	for i := 0; i < n; i++ {
		ans = max(ans, (right[i]-left[i]-1)*heights[i])
	}
	return ans
}

func largestRectangleArea1(heights []int) int {
	var max = func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}

	n := len(heights)
	var stack []int
	ans := 0
	for i := 0; i < n; i++ {
		for len(stack) > 0 && heights[stack[len(stack)-1]] >= heights[i] {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			left := -1
			if len(stack) > 0 {
				left = stack[len(stack)-1]
			}
			ans = max(ans, (i-left-1)*heights[top])
		}
		stack = append(stack, i)
	}
	fmt.Printf("len:%v\n", len(stack))
	slen := len(stack)
	for i := 0; i < slen; i++ {
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		left := -1
		if len(stack) > 0 {
			left = stack[len(stack)-1]
		}
		ans = max(ans, (n-left-1)*heights[top])
	}
	return ans
}

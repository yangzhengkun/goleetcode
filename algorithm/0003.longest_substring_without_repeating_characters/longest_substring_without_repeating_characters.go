package _003_longest_substring_without_repeating_characters

//给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。
//
//示例 1:
//
//输入: "abcabcbb"
//输出: 3
//解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
//示例 2:
//
//输入: "bbbbb"
//输出: 1
//解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
//示例 3:
//
//输入: "pwwkew"
//输出: 3
//解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
//     请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。
//

// 双指针+滑动窗口
func lengthOfLongestSubstring(s string) int {
	var max = func(i, j int) int {
		if i > j {
			return i
		} else {
			return j
		}
	}
	var freq = make([]int, 256)
	l, r, res := 0, -1, 0
	for l < len(s) {
		// 循环变量l和r的含义: s[l...r]之间不包含重复字母
		if r+1 < len(s) && freq[s[r+1]] == 0 {
			r++
			freq[s[r]] ++
		} else {
			freq[s[l]]--
			l++
		}
		res = max(res, r-l+1)
	}
	return res
}

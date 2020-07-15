package _125_valid_palindrome

import (
	"strings"
	"unicode"
)

//给定一个字符串，验证它是否是回文串，只考虑字母和数字字符，可以忽略字母的大小写。
//
//说明：本题中，我们将空字符串定义为有效的回文串。
//
//示例 1:
//
//输入: "A man, a plan, a canal: Panama"
//输出: true
//示例 2:
//
//输入: "race a car"
//输出: false
//

//指针对撞
func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	left, right := 0, len(s)-1
	for left < right {
		for left < right && !isalnum(s[left]) {
			left++
		}
		for left < right && !isalnum(s[right]) {
			right--
		}
		if left < right {
			if s[left] != s[right] {
				return false
			}
			left++
			right--
		}
	}
	return true
}

func isalnum(ch byte) bool {
	return (ch >= 'A' && ch <= 'Z') || (ch >= 'a' && ch <= 'z') || (ch >= '0' && ch <= '9')
}

//使用unicode包判断
func isPalindrome1(s string) bool {
	//使用unicode包判断是否为字母或数字
	isValid := func(char rune) bool {
		return unicode.IsLetter(char) || unicode.IsDigit(char)
	}
	left, right := 0, len(s)-1
	for left < right {
		cLeft, cRight := rune(s[left]), rune(s[right])
		if !isValid(cLeft) && !isValid(cRight) {
			//两端都不是字母或数字
			left++
			right--
		} else if !isValid(cLeft) {
			// 左端不是字母或数字
			left++
		} else if !isValid(cRight) {
			right--
		} else if unicode.ToLower(cLeft) != unicode.ToLower(cRight) {
			//两端都是数字或字母,但是不相等
			return false
		} else {
			//两端都是数字或字母,且相等
			left++
			right--
		}
	}
	return true
}

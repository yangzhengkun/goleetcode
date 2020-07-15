package _345_reverse_vowels_of_a_string

//编写一个函数，以字符串作为输入，反转该字符串中的元音字母。
//
//示例 1:
//
//输入: "hello"
//输出: "holle"
//示例 2:
//
//输入: "leetcode"
//输出: "leotcede"
//说明:
//元音字母不包含字母"y"。
//

func reverseVowels(s string) string {
	if s == "" {
		return s
	}
	isVowel := func(char byte) bool {
		switch char {
		case 'A', 'a', 'E', 'e', 'I', 'i', 'O', 'o', 'U', 'u':
			return true
		}
		return false
	}

	res := []byte(s)
	left, right := 0, len(res)-1
	for left < right {
		if !isVowel(res[left]) && !isVowel(res[right]) {
			//两端都不是元音字母,相向而行
			left++
			right--
		} else if !isVowel(res[left]) {
			//左端不是元音字母,向右移动一个
			left++
		} else if !isVowel(res[right]) {
			// 右端不是元音字母,向左移动一个
			right--
		} else {
			//两端都是元音字母,则交换一下
			res[left], res[right] = res[right], res[left]
			left++
			right--
		}
	}
	return string(res)
}

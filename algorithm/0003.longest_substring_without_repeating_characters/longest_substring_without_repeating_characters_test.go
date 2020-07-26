package _003_longest_substring_without_repeating_characters

import (
	"fmt"
	"testing"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	s := "abcabcbb"
	fmt.Println(lengthOfLongestSubstring(s))
}

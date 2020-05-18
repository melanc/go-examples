package other

import (
	"fmt"
	"testing"
)

func TestLongestPalindrome2(t *testing.T){
	words := "cabbad"
	res := longestPalindrome2(words)
	fmt.Println(res)
}

func longestPalindrome2(s string) string {
	if len(s) <= 1 {
		return s
	}
	maxLen := 0
	start, end := 0, 0
	for i:=0; i<len(s); i++ {
		len1 := centerExpand(s, i, i)
		len2 := centerExpand(s, i, i+1)
		fmt.Println(i, len1, len2)
		if len1 > len2 {
			if maxLen < len1 {
				maxLen = len1
				start = i - len1/2
				end = i + len1/2
			}
		} else {
			if maxLen < len2 {
				maxLen = len2
				start = i - len2/2 + 1
				end = i + len2/2
			}
		}
	}
	return s[start:end+1]
}

func centerExpand(s string, c1 int, c2 int) int {
	if c1 < 0 || c2 > len(s)-1{
		return 0
	}
	l := len(s)
	left, right := c1, c2
	for left >= 0 && right < l {
		if s[left] != s[right] {
			break
		}
		left--
		right++
	}
	return right - left - 1
}
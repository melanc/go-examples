package other

import (
	"fmt"
	"testing"
)

func TestLongestPalindrome(t *testing.T){
	words := "civilwartestingwhetherthatnaptionoranynartionsoconceivedandsodedicatedcanlongendureWeareqmetonagreatbattlefiemldoftzhatwarWehavecometodedicpateaportionofthatfieldasafinalrestingplaceforthosewhoheregavetheirlivesthatthatnationmightliveItisaltogetherfangandproperthatweshoulddothisButinalargersensewecannotdedicatewecannotconsecratewecannothallowthisgroundThebravelmenlivinganddeadwhostruggledherehaveconsecrateditfaraboveourpoorponwertoaddordetractTgheworldadswfilllittlenotlenorlongrememberwhatwesayherebutitcanneverforgetwhattheydidhereItisforusthelivingrathertobededicatedheretotheulnfinishedworkwhichtheywhofoughtherehavethusfarsonoblyadvancedItisratherforustobeherededicatedtothegreattdafskremainingbeforeusthatfromthesehonoreddeadwetakeincreaseddevotiontothatcauseforwhichtheygavethelastpfullmeasureofdevotionthatweherehighlyresolvethatthesedeadshallnothavediedinvainthatthisnationunsderGodshallhaveanewbirthoffreedomandthatgovernmentofthepeoplebythepeopleforthepeopleshallnotperishfromtheearth"
	res := longestPalindrome(words)
	fmt.Println(res)
}

func longestPalindrome(s string) string {
	if len(s) <= 1 {
		return s
	}
	dp := make([]int, len(s))
	for i:=0; i< len(s); i++ {
		dp[i] = 1
	}
	for i:=0; i<len(s); i++ {
		ts := s[i:i+1]
		for j:=i+1; j<len(s); j++ {
			ts =  ts + s[j:j+1]
			if isPalindrome(ts){
				dp[i] = len(ts)
			}
		}
	}
	maxIdx := maxSlice1(dp)
	return s[maxIdx:maxIdx+dp[maxIdx]]
}

func maxSlice1(arr []int) int {
	max := 0
	for i:=0; i<len(arr); i++ {
		if arr[i] > arr[max] {
			max = i
		}
	}
	return max
}

func isPalindrome(s string) bool {
	l := len(s)
	if l == 0 {
		return true
	}
	for i:=0; i<l/2; i++ {
		if s[i] != s[l-i-1] {
			return false
		}
	}
	return true
}

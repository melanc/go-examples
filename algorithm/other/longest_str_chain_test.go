package other

import (
	"fmt"
	"sort"
	"testing"
)

func TestLongestStrChain(t *testing.T) {
	words := []string{"ft","t","q","ni","vw","oq","qq","ooq","ddp","vcw","nci","vft","eik","dqq","njci","ddpj","dqqt","eihk","oopq","oopwq","njpci","ddgpj","eihzk","doqqt","eihzke","oopwqq","wdoqqt","ddgpjy","eiihzke","wdeoqqt","wdkeoqqt","ebiihzke","iwdkeoqqt","iwdkeoqqtd"}
	res := longestStrChain(words)
	fmt.Println("res = ", res)
}

func longestStrChain(words []string) int {
	if len(words) == 0 {
		return 0
	}
	dp := make(map[string]int)
	for i:=0; i<len(words); i++ {
		dp[words[i]] = 1
	}
	sort.Sort(strSlice(words))
	fmt.Println(words)
	for i:=0; i<len(words); i++ {
		if i>=1 && words[i] == words[i-1] {
			continue
		}
		for j:=i+1; j<len(words); j++ {
			if len(words[j]) == len(words[i]) {
				continue
			}
			if len(words[j]) - 1 > len(words[i]) {
				break
			}
			if len(words[j]) - 1 == len(words[i]) && find(words[j], words[i]) {
				dp[words[j]] = max(dp[words[j]], dp[words[i]] + 1)
			}
		}
	}
	return maxSlice(dp)
}

func find(str string, s string) bool {
	for i:=0; i<len(str); i++ {
		t := str[0:i] + str[i+1:]
		if t == s {
			return true
		}
	}
	return false
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func maxSlice(arr map[string]int) int {
	if len(arr) == 0{
		return 0
	}
	max := 0
	for _, v := range arr {
		if v > max {
			max = v
		}
	}
	return max
}

type strSlice []string

func (s strSlice) Len() int {
	return len(s)
}

func (s strSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s strSlice) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}
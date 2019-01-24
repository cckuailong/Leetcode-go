package main

import (
	"fmt"
)

func lengthOfLongestSubstring(s string) int {
	var ok bool
	var ans int
	var c byte
	mm := map[byte]int{}
	ans = 0
	j := -1
	for i:=0;i < len(s);i++{
		c = s[i]
		_, ok = mm[c]    // has exist or not
		if ok{
			if mm[c] > j{   // keep the slide right forward
				j = mm[c]
			}
		}
		if i-j > ans{
			ans = i-j
		}
		mm[c] = i
	}
	return ans
}

func main(){
	s := "abba"
	fmt.Print(lengthOfLongestSubstring(s))
}

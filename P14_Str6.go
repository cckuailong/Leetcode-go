package main

import "fmt"

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0{
		return ""
	}
	res := ""
	for i:=0;i<len(strs[0]);i++{
		first := strs[0][i]
		for _, str := range(strs){
			if i >= len(str) || first != str[i]{
				return res
			}
		}
		res += string(first)
	}
	return res
}

func main(){
	strs := []string{"aa", "a"}
	fmt.Println(longestCommonPrefix(strs))
}

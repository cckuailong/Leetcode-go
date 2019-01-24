package main

import (
	"fmt"
	"regexp"
)

func isMatch(s string, p string) bool {
	re := regexp.MustCompile(p)
	res := re.FindStringSubmatch(s)
	if res == nil{
		return false
	}else{
		if res[0] == s{
			return true
		}else{
			return false
		}
	}
}

func main(){
	s := "aa"
	p := "a"
	fmt.Print(isMatch(s, p))
}

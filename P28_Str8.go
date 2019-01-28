package main

import (
	"fmt"
	"strings"
)

func strStr(haystack string, needle string) int {
	return strings.Index(haystack, needle)
}

func main(){
	a := "aaaaaa"
	b := ""
	fmt.Println(strStr(a,b))
}


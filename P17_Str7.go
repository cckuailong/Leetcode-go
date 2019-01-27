package main

import "fmt"

var mappings = make(map[int]string)

func init(){
	mappings[2] = "abc"
	mappings[3] = "def"
	mappings[4] = "ghi"
	mappings[5] = "jkl"
	mappings[6] = "mno"
	mappings[7] = "pqrs"
	mappings[8] = "tuv"
	mappings[9] = "wxyz"
}

func letterCombinations(digits string) []string {
	res := []string{}
	if len(digits) == 0{
		return res
	}
	backtrack("", digits, &res)
	return res
}

func backtrack(s, digits string, res *[]string){
	if (len(digits) == 0){
		*res = append(*res, s)
		return
	}
	for _,c := range(mappings[int(digits[0]-'0')]){
		backtrack(s+string(c), digits[1:], res)
	}
}

func main(){
	test := ""
	fmt.Println(letterCombinations(test))
}

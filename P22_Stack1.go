package main

import "fmt"

func generateParenthesis(n int) []string {
	res := []string{}
	dfs("", 0, 0, n, &res)
	return res
}

func dfs(str string, left, right, n int, res *[]string){
	if (right > left || left > n || right > n){
		return
	}
	if len(str) == 2*n{
		*res = append(*res, str)
		return
	}
	dfs(str+"(", left+1, right, n, res)
	dfs(str+")", left, right+1, n, res)
}

func main(){
	res := generateParenthesis(3)
	fmt.Print(res)
}

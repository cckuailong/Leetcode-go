package main

import (
	"fmt"
	"math"
)

func convert(s string, numRows int) string {
	res := ""
	n := len(s)
	buck := make([]string, n)
	j := 0
	forward := false
	if numRows == 1{
		return s
	}
	for i:=0;i<n;i++{
		buck[j] += string(s[i])
		if j == 0 || j == numRows-1{
			forward = !forward
		}
		if forward{
			j++
		}else{
			j--
		}
	}
	nn := int(math.Min(float64(n), float64(numRows)))
	for i:=0;i<nn;i++{
		res += buck[i]
	}
	return res
}

func main(){
	s := "LEETCODEISHIRING"
	numRows := 3
	fmt.Print(convert(s, numRows))
}

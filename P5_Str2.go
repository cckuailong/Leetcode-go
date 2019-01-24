package main

import (
	"fmt"
	"math"
)

func longestPalindrome(s string) string {
	// generate the Len
	new_s := []byte{'^', '#'}
	for _, c := range(s){
		new_s = append(new_s, byte(c))
		new_s = append(new_s, '#')
	}
	new_s = append(new_s, '$')

	// cal Manacher
	mx := 0
	po := 0
	res := 0
	start := 0
	n := len(new_s)
	Len := make([]int, n)
	for i:=1;i < n-1;i++{
		if mx > i{
			Len[i] = int(math.Min(float64(mx-i), float64(Len[2*po-i])))
		}else{
			Len[i] = 0
		}
		for new_s[i-1-Len[i]] == new_s[i+1+Len[i]]{
			Len[i]++
		}
		if Len[i]+i > mx{
			mx = Len[i] + i
			po = i
		}
		if Len[i] > res{
			res = Len[i]
			start = i
		}
	}
	beg := (start-1-res)/2
	end := beg + res
	return string(s[beg:end])
}

func main(){
	s := "babad"
	fmt.Print(longestPalindrome(s))
}

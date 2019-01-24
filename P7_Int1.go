package main

import (
	"fmt"
	"math"
)

func reverse(x int) int {
	res := 0
	for x != 0{
		tmp := x % 10
		x /= 10
		res = res * 10 + tmp
		if res > math.MaxInt32 || res < math.MinInt32{
			return 0
		}
	}
	return res
}

func main(){
	test := 1534236469
	fmt.Print(reverse(test))
}

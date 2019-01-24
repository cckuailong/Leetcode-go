package main

import (
	"fmt"
	"math"
)

func maxArea(height []int) int {
	max := 0.0
	l := 0
	r := len(height) - 1
	for l < r{
		max = math.Max(max, math.Min(float64(height[l]), float64(height[r])) * float64(r-l))
		if height[l] < height[r]{
			l++
		}else{
			r--
		}
	}
	return int(max)
}

func main(){
	height := []int{1,8,6,2,5,4,8,3,7}
	fmt.Print(maxArea(height))
}

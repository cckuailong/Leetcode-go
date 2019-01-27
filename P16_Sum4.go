package main

import (
	"fmt"
	"math"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	var res, dist, diff int
	sort.Ints(nums)
	diff = math.MaxInt64
	for i,_ := range(nums){
		l := i+1
		r := len(nums)-1
		if i == r{
			break
		}
		for(l<r){
			sum := nums[i] + nums[l] + nums[r]
			if sum < target{
				dist = target - sum
				if dist < diff{
					res = sum
					diff = dist
				}
				l++
			}else if(sum > target){
				dist = sum - target
				if dist < diff{
					res = sum
					diff = dist
				}
				r--
			}else{
				return sum
			}
		}
	}
	return res
}

func main(){
	nums := []int{-1, 2, 1, -4}
	target := 1
	fmt.Println(threeSumClosest(nums, target))
}

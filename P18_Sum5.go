package main

import (
	"fmt"
	"sort"
)

func fourSum(nums []int, target int) [][]int {
	res := [][]int{}
	ll := len(nums)
	sort.Ints(nums)
	for i:=0;i<ll-3;i++{
		for j:=i+1;j<ll-2;j++{
			l := j+1
			r := ll-1
			for(l<r){
				sum := nums[i] + nums[j] + nums[l] + nums[r]
				if sum < target{
					l++
				}else if(sum > target){
					r--
				}else{
					tmp := []int{nums[i], nums[j], nums[l], nums[r]}
					if !isIn(res, tmp){
						res = append(res, tmp)
					}
					l++
				}
			}
		}
	}
	return res
}

func isIn(list [][]int, item []int) bool{
	for _, cc := range(list){
		if isEqual(cc, item){
			return true
		}
	}
	return false
}

func isEqual(a,b []int)bool{
	l := len(a)
	for i:=0;i<l;i++{
		if a[i] != b[i]{
			return false
		}
	}
	return true
}

func main(){
	nums := []int{-3,-2,-1,0,0,1,2,3}
	target := 0
	fmt.Println(fourSum(nums, target))
}

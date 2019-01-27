package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	var tmp int
	var items [][]int
	list := [][]int{}
	l := len(nums)
	if l < 3{
		return list
	}
	if l == 3{
		sum := 0
		for i:=0;i<3;i++{
			sum += nums[i]
		}
		if sum == 0{
			return append(list,nums)
		}else{
			return list
		}
	}
	sort.Ints(nums)
	for i, c :=range(nums){
		tmp = -1*c
		tmp_l := make([]int,l, l)
		copy(tmp_l, nums)
		tmp_l = append(tmp_l[:i], tmp_l[i+1:]...)
		items = twoSum1(tmp_l, tmp)
		if len(items) != 0{
			for _,item := range(items){
				item = append(item, c)
				sort.Ints(item)
				if !isIn(list, item){
					list = append(list, item)
				}
			}
		}
	}
	return list
}

func twoSum1(nums []int, target int) [][]int {
	list := [][]int{}
	mm := map[int]int{}
	for i, num := range(nums){
		c_num := target - num
		if _, ok := mm[c_num];ok == true{
			tmp := []int{c_num, nums[i]}
			if !isIn(list, tmp){
				list = append(list, tmp)
			}
		}
		mm[num] = i
	}
	return list
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
	nums := []int{-4,-2,-2,-2,0,1,2,2,2,3,3,4,4,6,6}
	fmt.Println(threeSum(nums))
}

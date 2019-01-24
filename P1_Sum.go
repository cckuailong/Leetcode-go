package main

import "fmt"

// 建立哈希，查找速度O(1)
func twoSum(nums []int, target int) []int {
	mm := map[int]int{}
	for i, num := range(nums){
		c_num := target - num
		if _, ok := mm[c_num];ok == true{
			return []int{mm[c_num], i}
		}
		mm[num] = i
	}
	return []int{}
}

func main(){
	nums := []int {2,7,11,15}
	target := 9
	fmt.Print(twoSum(nums, target))
}

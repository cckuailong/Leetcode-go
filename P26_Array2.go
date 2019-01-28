package main

import "fmt"

func removeDuplicates(nums []int) int {
	if len(nums) == 0{
		return 0
	}
	i := 0
	j := 1
	for j < len(nums){
		if nums[i] != nums[j]{
			i++
			nums[i] = nums[j]
		}
		j++
	}
	return i+1
}

func main(){
	nums := []int{}
	fmt.Println(removeDuplicates(nums))
	fmt.Println(nums)
}

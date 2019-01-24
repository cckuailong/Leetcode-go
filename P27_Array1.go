package main

import "fmt"

func removeElement(nums []int, val int) int {
	i := 0
	for _, c := range(nums){
		if c != val{
			nums[i] = c
			i++
		}
	}
	return i
}

func main(){
	nums := []int{0,1,2,2,3,0,4,2}
	val := 2
	len := removeElement(nums, val)
	for i:=0;i<len;i++{
		fmt.Println(nums[i])
	}
}

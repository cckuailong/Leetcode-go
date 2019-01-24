package main

import "fmt"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var pos int
	i := 0
	j := 0
	n := len(nums1) + len(nums2)
	res := []int{}
	for {
		if i >= len(nums1){
			res = append(res, nums2[j:]...)
			break
		}
		if j >= len(nums2){
			res = append(res, nums1[i:]...)
			break
		}

		if nums1[i] < nums2[j]{
			res = append(res, nums1[i])
			i++
		}else{
			res = append(res, nums2[j])
			j++
		}
	}

	pos = n / 2
	if n % 2 == 0{
		return float64(res[pos]+res[pos-1])/2
	}else{
		return float64(res[pos])
	}
}

func main(){
	nums1 := []int{1,2}
	nums2 := []int{3,4}
	fmt.Print(findMedianSortedArrays(nums1, nums2))
}

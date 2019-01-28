package main

import "fmt"

func divide(dividend int, divisor int) int {
	cnt, closest := 0, 0
	flag_cnt := 0
	if dividend == 0{
		return 0
	}else if dividend < 0{
		dividend = -dividend
		flag_cnt++
	}
	if divisor < 0{
		divisor = -divisor
		flag_cnt++
	}
	for i:=31;i>=0;i--{
		if closest + divisor<<uint(i) <= dividend{
			closest += divisor<<uint(i)
			cnt |= 1<<uint(i)
		}
	}
	if flag_cnt != 1{
		if cnt > 2147483647{
			return 2147483647
		}else{
			return cnt
		}
	}
	return -cnt
}

func main(){
	fmt.Println(divide(-2147483648, -1))
}


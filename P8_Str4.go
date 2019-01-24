package main

import (
	"fmt"
	"strings"
)

func myAtoi(str string) int {
	neg := false
	res := 0
	str = strings.TrimSpace(str)
	if len(str) == 0{
		return 0
	}
	if str[0] == '-'{
		if len(str) == 1{
			return 0
		}
		str = str[1:]
		neg = true
	}else if str[0] == '+'{
		if len(str) == 1{
			return 0
		}
		str = str[1:]
	}
	if str[0] > '9' || str[0] < '0'{
		return 0
	}
	for _, c := range(str){
		if c >= '0' && c <= '9'{
			if neg{
				if res > 214748364 || res == 214748364 && c >= '9'{
					return -2147483648
				}
			}else{
				if res > 214748364 || res == 214748364 && c >= '8'{
					return 2147483647
				}
			}
			res = res * 10 + int(c) - 48
		}else{
			break
		}
	}

	if neg{
		return (-1)*res
	}else{
		return res
	}
}

func main(){
	str := "9223372036854775807"
	fmt.Println(myAtoi(str))
}

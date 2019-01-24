package main

import "fmt"

func romanToInt(s string) int {
	num := 0
	s += " "
	for i:=0;i<len(s);i++{
		if string(s[i]) == " "{
			break
		}
		if string(s[i]) == "M"{
			num += 1000
		}else if string(s[i]) == "D"{
			num += 500
		}else if string(s[i]) == "C"{
			if string(s[i+1]) == "M"{
				num += 900
				i++
			}else if string(s[i+1]) == "D"{
				num += 400
				i++
			}else{
				num += 100
			}
		}else if string(s[i]) == "L"{
			num += 50
		}else if string(s[i]) == "X"{
			if string(s[i+1]) == "C"{
				num += 90
				i++
			}else if string(s[i+1]) == "L"{
				num += 40
				i++
			}else{
				num += 10
			}
		}else if string(s[i]) == "V"{
			num += 5
		}else if string(s[i]) == "I"{
			if string(s[i+1]) == "X"{
				num += 9
				i++
			}else if string(s[i+1]) == "V"{
				num += 4
				i++
			}else{
				num += 1
			}
		}
	}
	return num
}

func main(){
	fmt.Println(romanToInt("MCMXCIV"))
}

package main

import "fmt"

func intToRoman(num int) string {
	var tmp string
	str := ""
	num, tmp = calRome(num, 1000, 900, "M", "CM")
	str += tmp
	num, tmp = calRome(num, 500, 400, "D", "CD")
	str += tmp
	num, tmp = calRome(num, 100, 90, "C", "XC")
	str += tmp
	num, tmp = calRome(num, 50, 40, "L", "XL")
	str += tmp
	num, tmp = calRome(num, 10, 9, "X", "IX")
	str += tmp
	num, tmp = calRome(num, 5, 4, "V", "IV")
	str += tmp
	for i:=0;i<num;i++{
		str += "I"
	}
	return str
}

func calRome(num, block, block2 int, sym, sym2 string)(int, string){
	str := ""
	if n:=num/block;n!=0{
		for i:=0;i<n;i++{
			str += sym
		}
	}
	num %= block
	if num / block2 != 0{
		str += sym2
	}
	num %= block2
	return num, str
}

func main(){
	fmt.Println(intToRoman(58))
}

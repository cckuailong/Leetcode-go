package main

func isPalindrome(x int) bool {
	res := 0
	x_pre := x
	if x < 0{
		return false
	}
	for x!=0{
		tmp := x % 10
		x /= 10
		res = res * 10 + tmp
	}
	return res == x_pre
}

func main(){

}

package main

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode{
	h := ListNode{0,nil}
	p := head
	for p!=nil{
		node := ListNode{p.Val, h.Next}
		h.Next = &node
		p = p.Next
	}
	return &h
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	h := reverseList(head)
	pre := h
	p := h.Next
	for cnt:=0;cnt < n-1;cnt++{
		pre = p
		p = p.Next
	}
	pre.Next = p.Next
	res := reverseList(h.Next)
	return res.Next
}

func (list *ListNode)append(num int){
	node := ListNode{num, list.Next}
	list.Next = &node
}

func main(){
	l1 := ListNode{1,nil}
	l1.append(4)
	l1.append(2)
	h := removeNthFromEnd(&l1,2)
	p := h
	for p!= nil{
		fmt.Println(p.Val)
		p = p.Next
	}
}

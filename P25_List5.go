package main

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	cnt := k
	h := ListNode{0, head}
	p := &h
	for p != nil{
		if cnt > k-1{
			swapOnePair(p, k)
			cnt = 0
		}
		cnt++
		p = p.Next
	}
	return h.Next
}

func swapOnePair(head *ListNode, k int){
	i := 0
	var firstnode *ListNode
	if head == nil{
		return
	}
	h := ListNode{0,nil}
	p := head.Next
	if p == nil {
		return
	}
	for i<k && p!= nil{
		node := ListNode{p.Val, h.Next}
		h.Next = &node
		if i == 0{
			firstnode = &node
		}
		p = p.Next
		i++
	}
	if p!= nil || p==nil && i==k{
		firstnode.Next = p
		head.Next = h.Next
	}
}

func (list *ListNode)append(num int){
	node := ListNode{num, list.Next}
	list.Next = &node
}

func main(){
	l1 := ListNode{1,nil}
	//l1.append(5)
	//l1.append(4)
	//l1.append(3)
	//l1.append(2)
	p := reverseKGroup(&l1, 3)
	for p != nil{
		fmt.Println(p.Val)
		p = p.Next
	}
}

package main

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	cnt := 2
	h := ListNode{0, head}
	p := &h
	for p != nil{
		if cnt > 1{
			swapOnePair(p)
			cnt = 0
		}
		cnt++
		p = p.Next
	}
	return h.Next
}

func swapOnePair(head *ListNode){
	firstNode := head.Next
	if firstNode == nil{
		return
	}
	secondNode := firstNode.Next
	if secondNode != nil{
		firstNode.Next = secondNode.Next
		secondNode.Next = firstNode
		head.Next = secondNode
	}
}

func (list *ListNode)append(num int){
	node := ListNode{num, list.Next}
	list.Next = &node
}

func main(){
	l1 := ListNode{1,nil}
	l1.append(4)
	l1.append(3)
	l1.append(2)
	p := swapPairs(&l1)
	for p != nil{
		fmt.Println(p.Val)
		p = p.Next
	}
}

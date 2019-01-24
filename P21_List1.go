package main

import (
	"fmt"
)

type ListNode struct {
	Val int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	res := ListNode{0, nil}
	cur := &res
	var p, q *ListNode
	p = l1
	q = l2
	for p != nil && q != nil{
		if p.Val < q.Val{
			cur.Next = p
			cur = cur.Next
			p = p.Next
		}else{
			cur.Next = q
			cur = cur.Next
			q = q.Next
		}
	}
	if p == nil{
		cur.Next = q
	}else{
		cur.Next = p
	}
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
	l2 := ListNode{1,nil}
	l2.append(4)
	l2.append(3)
	res := mergeTwoLists(&l1, &l2)
	p := res
	for p != nil{
		fmt.Println(p.Val)
		p = p.Next
	}
}

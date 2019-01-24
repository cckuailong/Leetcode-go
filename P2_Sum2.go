package main

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var p1_num, p2_num int
	var sum, c int
	res_node := ListNode{0, nil}
	curr := &res_node
	p1 := l1
	p2 := l2
	c = 0
	for p1 != nil || p2 != nil{
		if p1 == nil{
			p1_num = 0
		}else{
			p1_num = p1.Val
			p1 = p1.Next
		}
		if p2 == nil{
			p2_num = 0
		}else{
			p2_num = p2.Val
			p2 = p2.Next
		}

		sum = p1_num + p2_num + c
		if sum > 9{
			sum = sum % 10
			c = 1
		}else{
			c = 0
		}
		new_node := ListNode{sum, nil}
		curr.Next = &new_node
		curr = curr.Next

	}
	if c == 1{
		new_node := ListNode{1,nil}
		curr.Next = &new_node
	}

	return res_node.Next
}

func (list *ListNode)append(num int){
	node := ListNode{num, list.Next}
	list.Next = &node
}

func main(){
	l1 := ListNode{2, nil}
	l1.append(3)
	l1.append(4)
	l2 := ListNode{5, nil}
	l2.append(4)
	l2.append(6)

	res_node := addTwoNumbers(&l1, &l2)
	p := res_node
	for p != nil{
		fmt.Print(p.Val)
		p = p.Next
	}
}

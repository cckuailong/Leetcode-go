package main

func mergeKLists(lists []*ListNode) *ListNode {
	ll := len(lists)
	interval := 1
	for interval < ll{
		for i:=0;i<ll-interval;i+=interval*2{
			lists[i] = mergeTwoLists(lists[i], lists[i+interval])
		}
		interval*=2
	}
	if ll > 0{
		return lists[0]
	}else {
		return nil
	}

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

func main(){

}

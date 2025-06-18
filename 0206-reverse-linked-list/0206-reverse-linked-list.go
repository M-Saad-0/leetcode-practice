package main

// type ListNode struct {
// 	Val  int
// 	Next *ListNode
// }

func reverseList(head *ListNode) *ListNode {
	copy := head
	copy1 := head
	var vals []int
	for copy != nil {
		vals = append(vals, copy.Val)
		copy = copy.Next
	}

	i := len(vals)-1
	for copy1 != nil {
		copy1.Val = vals[i]
		i--
        copy1 = copy1.Next
	}
	return head
}

package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(tmp1 *ListNode, tmp2 *ListNode) *ListNode {

	dummy := &ListNode{}
	result := dummy
	for tmp1 != nil && tmp2 != nil {
		if tmp1.Val < tmp2.Val {
			result.Next = tmp1
			tmp1 = tmp1.Next
		} else {
			result.Next = tmp2
			tmp2 = tmp2.Next
		}
		result = result.Next
	}

	if tmp1 != nil {
		result.Next = tmp1
	} else {
		result.Next = tmp2
	}
	return dummy.Next
}

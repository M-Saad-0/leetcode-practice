/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    list1 := []int{}
    list2 := []int{}

    for l1 != nil {
        list1 = append(list1, l1.Val)
        l1 = l1.Next
    }
    for l2 != nil {
        list2 = append(list2, l2.Val)
        l2 = l2.Next
    }

    n := len(list1)
    m := len(list2)
    i, j := 0, 0
    carry := 0
    res := []int{}

    for i < n || j < m || carry > 0 {
        sum := carry
        if i < n {
            sum += list1[i]
            i++
        }
        if j < m {
            sum += list2[j]
            j++
        }
        res = append(res, sum%10)
        carry = sum / 10
    }

    // Build the linked list from result
    dummy := &ListNode{}
    curr := dummy
    for _, val := range res {
        curr.Next = &ListNode{Val: val}
        curr = curr.Next
    }

    return dummy.Next
}

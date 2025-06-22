/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func middleNode(head *ListNode) *ListNode {
    var l []int = []int{}
    for head != nil {
        l=append(l, head.Val)
        head = head.Next
    }
    fmt.Println(l)
    midL := []int{}
    if len(l)%2 == 0 {
        midL = l[(len(l)/2):len(l)]
    } else {
        midL = l[(len(l)/2):len(l)]
    }
    fmt.Println(midL)


    dummy := &ListNode{}
    head1 := dummy
    for i:=0; i<len(midL); i++  {
        head1.Next = &ListNode{Val: midL[i]}
        head1 = head1.Next 
    }
    return dummy.Next
}
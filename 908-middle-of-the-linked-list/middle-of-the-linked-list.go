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
    
    
  


    dummy := &ListNode{}
    head1 := dummy
    for i:=len(l)/2; i<len(l); i++  {
        head1.Next = &ListNode{Val: l[i]}
        head1 = head1.Next 
    }
    return dummy.Next
}
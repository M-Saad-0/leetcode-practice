// /**
//  * Definition for singly-linked list.
//  * type ListNode struct {
//  *     Val int
//  *     Next *ListNode
//  * }
//  */
// func middleNode(head *ListNode) *ListNode {
//     var l []int = []int{}
//     for head != nil {
//         l=append(l, head.Val)
//         head = head.Next
//     }
    
//     l = l[len(l)/2:len(l)]
//     var dummy *ListNode = &ListNode{}
//     head1 := dummy
//     for i:=0; i<len(l); i++  {
//         head1.Next = &ListNode{Val: l[i]}
//         head1 = head1.Next 
//     }
//     return dummy.Next
// }



func middleNode(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }
    slow := head
    fast := head
    for fast != nil && fast.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next
    }
    return slow
}
package main


/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
    var slc map[*ListNode]int = map[*ListNode]int{}
    for head != nil {
        _, exists := slc[head]
        if exists {
            return true
        } else {
            slc[head] = head.Val
        }
        head = head.Next 
    }
    return false
}
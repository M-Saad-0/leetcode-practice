/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSymmetric(root *TreeNode) bool {
    if root == nil {
        return true
    }

    right := root.Right
    left := root.Left

    r := []int{}
    l := []int{}
    toRList(right, &r)
    toLList(left, &l)
    if len(r)!=len(l) {
        return false
    }
        fmt.Println(l)
        fmt.Println(r)

    for i:=0;i<len(r);i++{
        if r[i]!=l[i] {
            return false
        }
    }
    return true

}

func toRList(root *TreeNode, lst *[]int) {
    if root == nil {
        *lst = append(*lst, -1)
        return
    }

    *lst = append(*lst, root.Val)
    toRList(root.Right, lst)
    toRList(root.Left, lst)


}

func toLList(root *TreeNode, lst *[]int) {
    if root == nil {
        *lst = append(*lst, -1)
        return
    }

    *lst = append(*lst, root.Val)
    toLList(root.Left, lst)
    toLList(root.Right, lst)


}
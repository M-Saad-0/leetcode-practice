package main
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func invertTree(root *TreeNode) *TreeNode {
    if(root==nil){
        return nil
    }
        swap(&root.Left, &root.Right)
    if(root.Left!=nil) {
        invertTree(root.Left)
    }
    if(root.Right!=nil) {
        invertTree(root.Right)
    }
    return root
 }

func swap(a **TreeNode, b **TreeNode) {
    tmp := *a
    *a = *b
    *b = tmp
}
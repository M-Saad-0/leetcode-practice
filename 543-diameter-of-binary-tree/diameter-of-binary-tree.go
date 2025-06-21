/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func diameterOfBinaryTree(root *TreeNode) int {
    var maxD int
    var dfs func(*TreeNode) int
    dfs = func(node *TreeNode) int {
        if node == nil {
        return 0
    } 
    
    left := dfs(node.Left)
    right := dfs(node.Right)
    
    if right + left > maxD {
        maxD = right + left
    }
    fmt.Printf("%d, %d\n", right, left)
    return 1 + max(right, left) 
    }


    dfs(root)
    return maxD
}

func max(a int, b int) int {
    if a > b {
        return a
    } else {
        return b
    }
}
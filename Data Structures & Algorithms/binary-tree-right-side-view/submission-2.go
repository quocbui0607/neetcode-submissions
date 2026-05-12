/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func rightSideView(root *TreeNode) []int {
    if root == nil {
        return nil
    }


    res := []int{}
    queue := []*TreeNode{root}

    for len(queue) > 0 {
        level := len(queue)
        for i := 0; i < level; i++ {
            node := queue[0]
            queue = queue[1:]
            if i == level -1 {
                res = append(res, node.Val)
            }

            if node.Left != nil {
                queue = append(queue, node.Left)
            }

             if node.Right != nil {
                queue = append(queue, node.Right)
            }
        }
    }

     return res
}

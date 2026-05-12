/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func levelOrder(root *TreeNode) [][]int {
    if root == nil {
        return nil
    }

    queue := []*TreeNode{root}
    res := [][]int{}
    for len(queue) > 0 {
        size := len(queue)
        levelNodes := []int{}
        for i := 0; i < size; i++ {
            node := queue[0]
            levelNodes = append(levelNodes, node.Val)
            queue = queue[1:]

            if node.Left != nil {
                queue = append(queue, node.Left)
            }

            if node.Right != nil {
                queue = append(queue, node.Right)
            }
        }

        res = append(res, levelNodes)
    }

    return res
}

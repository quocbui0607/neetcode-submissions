/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isSameTree(p *TreeNode, q *TreeNode) bool {
     var dfs func(node1 *TreeNode, node2 *TreeNode) bool

	dfs = func(node1, node2 *TreeNode) bool {
		if node1 == nil && node2 == nil  {
			return true
		}

        if node1 == nil || node2 == nil  {
			return false
		}

		isSameLeft := dfs(node1.Left, node2.Left)
		if !isSameLeft {
			return false
		}

		isSameRight := dfs(node1.Right, node2.Right)
		if !isSameRight {
			return false
		}

		return isSameLeft && isSameRight && node1.Val == node2.Val
	}

	return dfs(p,q)
}

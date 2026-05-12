/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func maxDepth(root *TreeNode) int {
	var dfs func(root *TreeNode) int
	dfs = func(root *TreeNode) int {
		if root == nil {
			return 0 
		}

		return int(math.Max(float64(1 + dfs(root.Left)), float64(1 + dfs(root.Right))))
	}

	return dfs(root)
}

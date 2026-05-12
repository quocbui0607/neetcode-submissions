/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func diameterOfBinaryTree(root *TreeNode) int {
    var dfs func(root *TreeNode) int
	max := 0
	dfs = func(root *TreeNode) int {
		if root == nil {
			return 0
		}

		leftH := dfs(root.Left)
		rightH := dfs(root.Right)

		if leftH + rightH > max {
            max = leftH + rightH
        }

		if leftH > rightH {
			return leftH + 1
		}

		return rightH + 1
	}

	dfs(root)

	return max
}

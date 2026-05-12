/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isBalanced(root *TreeNode) bool {
    var dfs func(root *TreeNode) int
	dfs = func(root *TreeNode) int {
		if root == nil {
			return 0
		}

		numLeft := dfs(root.Left)
		if numLeft == -1 {
			return -1
		}
		numRight := dfs(root.Right)
		if numRight == -1 {
			return -1
		}

		isBalanced := int(math.Abs(float64(numLeft - numRight))) < 2
		if !isBalanced {
			return -1
		}

		if numLeft > numRight {
			return numLeft +1
		}


		return numRight + 1
		
	}

	return dfs(root) != -1
} 

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func hasPathSum(root *TreeNode, targetSum int) bool {
	acc := 0
	var calculateSum func(root *TreeNode) bool
	calculateSum = func(root *TreeNode) bool {
		if root == nil {
			return false
		}
		acc += root.Val
		if root.Left == nil && root.Right == nil {
			if acc == targetSum {
				return true
			}
			acc -= root.Val
			return false
		}

		if calculateSum(root.Left) {
			return true
		}

		if calculateSum(root.Right) {
			return true
		}

		acc -= root.Val
		return false
	}

	return calculateSum(root)

}

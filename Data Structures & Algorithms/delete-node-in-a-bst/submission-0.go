/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 
func findMin(root *TreeNode) *TreeNode {
	min := root

	for min != nil && min.Left != nil {
		min = min.Left
	}

	return min
}

func deleteNode(root *TreeNode, key int) *TreeNode {
	// in case not found
	if root == nil {
		return nil
	}

	if root.Val < key {
		root.Right = deleteNode(root.Right, key)
	} else if root.Val > key {
		root.Left = deleteNode(root.Left, key)
	} else {
		// we are in value that we want to remove
		if root.Left == nil {
			return root.Right
		} else if root.Right == nil {
			return root.Left
		} else {
			// has 2 children
			// find lowest 
			lowest := findMin(root.Right)
			// then replace 
			root.Val = lowest.Val
			// then cleanup - remove lowest because now it lives in another place
			root.Right = deleteNode(root.Right, lowest.Val)
		}	
	}

	return root
}

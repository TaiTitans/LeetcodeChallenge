package main


type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func insert(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val}
	}
	if val < root.Val {
		root.Left = insert(root.Left, val)
	} else {
		root.Right = insert(root.Right, val)
	}
	return root
}

func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}
	if key < root.Val {
		root.Left = deleteNode(root.Left, key)
	} else if key > root.Val {
		root.Right = deleteNode(root.Right, key)
	} else {
		if root.Left == nil {
			return root.Right
		} else if root.Right == nil {
			return root.Left
		}
		minNode := findMin(root.Right)
		root.Val = minNode.Val
		root.Right = deleteNode(root.Right, minNode.Val)
	}

	return root
}

func findMin(node *TreeNode) *TreeNode {
	current := node
	for current.Left != nil {
		current = current.Left
	}
	return current
}

func findMinDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftDepth := findMinDepth(root.Left)
	rightDepth := findMinDepth(root.Right)
	if leftDepth == 0 {
		return rightDepth + 1
	}
	if rightDepth == 0 {
		return leftDepth + 1
	}
	if leftDepth < rightDepth {
		return leftDepth + 1
	}
	return rightDepth + 1
}



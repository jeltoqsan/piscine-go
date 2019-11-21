package main

import "fmt"

type TreeNode struct {
	Left, Right, Parent *TreeNode
	Data                string
}

func BTreeLevelCount(root *TreeNode) int {
	x := 1
	y := 1
	if root.Left == nil && root.Right == nil {
		return 1
	}
	if root.Left != nil {
		x += BTreeLevelCount(root.Left)
	}
	if root.Right != nil {
		y += BTreeLevelCount(root.Right)
	}
	if x >= y {
		return x
	} else {
		return y
	}
}

func main() {
	root := &TreeNode{Data: "4"}
	BTreeInsertData(root, "1")
	BTreeInsertData(root, "7")
	BTreeInsertData(root, "5")
	fmt.Println(BTreeLevelCount(root))
}

func BTreeInsertData(root *TreeNode, data string) *TreeNode {
	if root == nil {
		return &TreeNode{Data: data}
	}
	if data < root.Data {
		root.Left = BTreeInsertData(root.Left, data)
	} else {
		root.Right = BTreeInsertData(root.Right, data)
	}
	return root
}

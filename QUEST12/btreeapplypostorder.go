package piscine

/*package main

import "fmt"

type TreeNode struct {
	Left, Right, Parent *TreeNode
	Data                string
}

func BTreeInsertData(root *TreeNode, data string) *TreeNode {
	if root == nil {
		return &TreeNode{Data: data}
	}
	if root.Data > data {
		root.Left = BTreeInsertData(root.Left, data)
	} else {
		root.Right = BTreeInsertData(root.Right, data)
	}
	if root.Left != nil {
		root.Left.Parent = root
	}
	if root.Right != nil {
		root.Right.Parent = root
	}
	return root
}*/
func BTreeApplyPostorder(root *TreeNode, f func(...interface{}) (int, error)) {
	if root == nil {
		return
	}
	BTreeApplyPostorder(root.Left, f)
	BTreeApplyPostorder(root.Right, f)
	f(root.Data)
}

/*func main() {
	root := &TreeNode{Data: "4"}
	BTreeInsertData(root, "1")
	BTreeInsertData(root, "7")
	BTreeInsertData(root, "5")
	BTreeApplyPostorder(root, fmt.Println)

}*/

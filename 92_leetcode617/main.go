package main

import "fmt"

func main() {

	root1 := new(TreeNode)
	root1.Val = 1
	root1.Left = new(TreeNode)
	root1.Left.Val = 3
	root1.Left.Left = new(TreeNode)
	root1.Left.Left.Val = 5

	root1.Right = new(TreeNode)
	root1.Right.Val = 2

	root2:=new(TreeNode)
	root2.Val=2
	root2.Left=new(TreeNode)
	root2.Left.Val=1
	root2.Left.Right=new(TreeNode)
	root2.Left.Right.Val=4

	root2.Right=new(TreeNode)
	root2.Right.Val=3
	root2.Right.Right=new(TreeNode)
	root2.Right.Right.Val=7
	head:=mergeTrees(root1,root2)
	fmt.Println(head.Val)

}

//Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 == nil && root2 == nil {
		return nil
	}
	head := &TreeNode{}
	if root1 != nil && root2 != nil {
		head.Val = root2.Val + root1.Val
	}
	if root2 != nil && root1 == nil {
		head.Val = root2.Val
	}
	if root1 != nil && root2 == nil {
		head.Val = root1.Val
	}
	if root1!=nil && root2!=nil{
		head.Right = mergeTrees(root1.Right, root2.Right)
	}
	if root1==nil && root2!=nil{
		head.Right = mergeTrees(nil, root2.Right)
	}
	if root2==nil && root1!=nil{
		head.Right = mergeTrees(root1.Right, nil)
	}

	if root1!=nil && root2!=nil{
		head.Left = mergeTrees(root1.Left, root2.Left)
	}
	if root1==nil && root2!=nil{
		head.Left = mergeTrees(nil, root2.Left)
	}
	if root2==nil && root1!=nil{
		head.Left = mergeTrees(root1.Left, nil)
	}

	return head
}

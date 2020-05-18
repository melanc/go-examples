package tree

import (
	"fmt"
	"testing"
)

type BiTree struct {
	Data   int
	LChild *BiTree
	RChild *BiTree
}

func TestBiTree(t *testing.T) {
	var a []int = []int{6, 1, 2, 5, 9, 3, 4, 7, 10, 8}
	tr := NewBiTree(a)
	PreOrder(tr)
	fmt.Println("")
	InOrder(tr)
	fmt.Println("")
	PostOrder(tr)
	fmt.Println("")

}

func NewBiTree(arr []int) *BiTree {
	var root *BiTree
	for _, v := range arr {
		root = InsertBiTree(root, v)
	}
	return root
}

func InsertBiTree(tree *BiTree, val int) *BiTree {
	t := &BiTree{
		Data:   val,
		LChild: nil,
		RChild: nil,
	}
	if tree == nil {
		return t
	} else {
		if (val<<2)%3 == 0 {
			tree.LChild = InsertBiTree(tree.LChild, val)
		} else {
			tree.RChild = InsertBiTree(tree.RChild, val)
		}
	}
	return tree
}

func PreOrder(tree *BiTree) {
	if tree != nil {
		fmt.Printf("%v ", tree.Data)
		PreOrder(tree.LChild)
		PreOrder(tree.RChild)
	}
}

func InOrder(tree *BiTree) {
	if tree != nil {
		PreOrder(tree.LChild)
		fmt.Printf("%v ", tree.Data)
		PreOrder(tree.RChild)
	}
}

func PostOrder(tree *BiTree) {
	if tree != nil {
		PreOrder(tree.LChild)
		PreOrder(tree.RChild)
		fmt.Printf("%v ", tree.Data)
	}
}

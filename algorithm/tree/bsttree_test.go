package tree

import (
	"fmt"
	"testing"
)

type BstTree struct {
	Data   int
	LChild *BstTree
	RChild *BstTree
}

func TestBSTTree(t *testing.T) {
	var a []int = []int{6, 1, 2, 5, 9, 3, 4, 7, 10, 8}
	tr := NewBST(a)
	PrintBST(tr)
	fmt.Println("")
	res := TraversalBST(tr)
	fmt.Println(res)
	b1 := SearchBST(tr, 5)
	fmt.Println(b1)
	b2 := SearchBST(tr, 15)
	fmt.Println(b2)

	tr2 := DeleteBST(tr, 4)
	PrintBST(tr2)
	fmt.Println("")
	tr4 := DeleteBST(tr, 9)
	PrintBST(tr4)
	fmt.Println("")
	tr3 := DeleteBST(tr, 10)
	PrintBST(tr3)
	fmt.Println("")
}

func NewBST(arr []int) *BstTree {
	var t *BstTree
	for _, v := range arr {
		t = InsertBST(t, v)
	}
	return t
}

func InsertBST(tree *BstTree, val int) *BstTree {
	if tree == nil {
		tree = &BstTree{
			Data:   val,
			LChild: nil,
			RChild: nil,
		}
	} else {
		if val < tree.Data {
			tree.LChild = InsertBST(tree.LChild, val)
		} else {
			tree.RChild = InsertBST(tree.RChild, val)
		}
	}
	return tree
}

func TraversalBST(tree *BstTree) []int {
	var res []int
	if tree != nil {
		r1 := TraversalBST(tree.LChild)
		res = append(res, r1...)
		res = append(res, tree.Data)
		r2 := TraversalBST(tree.RChild)
		res = append(res, r2...)
	}
	return res
}

// 中序遍历
func PrintBST(tree *BstTree) {
	if tree != nil {
		PrintBST(tree.LChild)
		fmt.Printf("%v ", tree.Data)
		PrintBST(tree.RChild)
	}
}

func SearchBST(tree *BstTree, val int) bool {
	if tree == nil {
		return false
	}
	if val == tree.Data {
		return true
	} else if val < tree.Data {
		return SearchBST(tree.LChild, val)
	} else {
		return SearchBST(tree.RChild, val)
	}
}

func DeleteBST(tree *BstTree, val int) *BstTree {
	if tree != nil {
		if val == tree.Data {
			tree = deleteBST(tree)
		} else if val < tree.Data {
			tree.LChild = DeleteBST(tree.LChild, val)
		} else {
			tree.RChild = DeleteBST(tree.RChild, val)
		}
	}
	return tree
}

func deleteBST(tree *BstTree) *BstTree {
	if tree.LChild == nil && tree.RChild == nil { // 该节点无子节点时
		return nil
	} else if tree.LChild == nil { // 只有右子节点
		rChild := tree.RChild
		tree.RChild = nil
		*tree = *rChild
	} else if tree.RChild == nil { // 只有左子节点
		lChild := tree.LChild
		tree.LChild = nil
		*tree = *lChild
	} else { // 同时有左右子节点
		lChild := tree.LChild
		rChild := tree.RChild
		tree.LChild = nil
		tree.RChild = nil
		// 左子节点代替被删除节点
		*tree = *lChild
		// 将被删除结点的右子树，放置于被删除结点的左子树的最右边
		t := lChild
		for t.RChild != nil {
			t = t.RChild
		}
		t.RChild = rChild
	}
	return tree
}

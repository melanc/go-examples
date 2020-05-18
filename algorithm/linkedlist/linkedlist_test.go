package linkedlist

import (
	"fmt"
	"testing"
)

type Node struct {
	Data int   //定义数据域
	Next *Node //定义下一个节点指针
}

func TestLinkedList(t *testing.T) {
	var a []int = []int{6, 1, 2, 5, 9, 3, 4, 7, 10, 8}
	l := MakeList(a)
	ShowList(l)
	fmt.Println(IsEmpty(l))
	fmt.Println(Length(l))
	Add(l, 12)
	ShowList(l)
	Append(l, 13)
	ShowList(l)
	l1 := Reverse(l)
	ShowList(l1)
}

func MakeList(arr []int) *Node {
	var head *Node
	var prev, tmp *Node
	for _, v := range arr {
		tmp = &Node{
			Data: v,
			Next: nil,
		}
		if head == nil {
			head = tmp
		}
		if prev != nil {
			prev.Next = tmp
		}
		prev = tmp

	}
	return head
}
func IsEmpty(l *Node) bool {
	if l == nil {
		return true
	} else {
		return false
	}
}

func Length(l *Node) int {
	curr := l
	cnt := 0
	for curr != nil {
		curr = curr.Next
		cnt++
	}
	return cnt
}

// 头部添加节点
func Add(l *Node, val int) {
	tmp := &Node{
		Data: val,
		Next: nil,
	}
	// 保存旧的头结点
	p := *l
	tmp.Next = &p
	// 指向新的头结点
	*l = *tmp
}

// 尾部添加节点
func Append(l *Node, val int) {
	tmp := &Node{
		Data: val,
		Next: nil,
	}
	if l == nil {
		l = tmp
	} else {
		curr := l
		for curr.Next != nil {
			curr = curr.Next
		}
		curr.Next = tmp
	}
}

func ShowList(l *Node) {
	curr := l
	for curr != nil {
		fmt.Printf("%v ", curr.Data)
		curr = curr.Next
	}
	fmt.Println("")
}

func Reverse(l *Node) *Node {
	// 三个临时变量
	var prev, curr, next *Node
	prev, curr, next = nil, l, nil
	if curr == nil || curr.Next == nil {
		return l
	}
	for curr != nil {
		next = curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	l.Next = nil
	return prev
}

package main

import (
	"fmt"
	"reflect"
)

//删除
func remove(slice []int, i int) []int {
	return append(slice[:i], slice[i+1:]...)
}

//新增
func add(slice []int, value int) []int {
	return append(slice, value)
}

//插入
func insert(slice *[]int, index int, value int) {
	rear := append([]int{}, (*slice)[index:]...)
	*slice = append(append((*slice)[:index], value), rear...)
}

//修改
func update(slice []int, index int, value int) {
	slice[index] = value
}

//查找
func find(slice []int, index int) int {
	return slice[index]
}

//清空slice
func empty(slice []int) []int{
	slice = append([]int{})
	return slice
}

//遍历
func list(slice []int) {
	for _, v := range slice {
		fmt.Printf("%d ", v)
	}
}

func main(){
	//数组，长度固定
	a := [3]int{1, 2}
	fmt.Println(a)
	// 数组，省略长度
	b := [...]int{1, 2, 3}
	fmt.Println(len(b))
	// 数组，指定索引
	c := [5]int{2: 50, 3: 100}
	fmt.Println(c)

	// 切片，没有长度
	s1 := []int{0, 1, 2, 3, 4, 5}
	fmt.Printf("s1 len: %d\n", cap(s1))
	fmt.Printf("s1 type: %T\n", reflect.TypeOf(s1))

	// 切片，使用make初始化
	s2 := make([]int, 6)
	fmt.Printf("s2 len: %d, cap: %d\n", len(s2), cap(s2))

	s2 = s1[:3]
	fmt.Printf("s2: %v, len: %d, cap: %d\n", s2, len(s2), cap(s2))

	s1 = add(s1, 6)
	fmt.Printf("s1: %v\n", s1)
	s1 = empty(s1)
	fmt.Printf("s1: %v\n", s1)

	for i:=0; i<10; i++ {
		s2 = add(s2, i)
	}
	fmt.Printf("s2: %v, len: %d, cap: %d\n", s2, len(s2), cap(s2))

	s3 := s2[3:8]
	fmt.Printf("s3: %v, len: %d, cap: %d\n", s2, len(s3), cap(s3))
	// 底层引用同一数组
	fmt.Printf("s2 ptr: %p, s3 ptr: %p\n", &s2[3], &s3[0])

}

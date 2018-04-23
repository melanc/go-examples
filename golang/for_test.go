package base

import (
	"fmt"
	"testing"
)

func TestFor(t *testing.T) {
	// 最基本的一种，单一条件循环
	// 这个可以代替其他语言的while循环
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	// 经典的循环条件初始化/条件判断/循环后条件变化
	for j := 1; j <= 5; j++ {
		fmt.Println(j)
	}

	// 无条件的for循环是死循环，除非你使用break跳出循环或者
	// 使用return从函数返回
	for {
		fmt.Println("loop")
		break
	}

	// 遍历切片
	x := []string{"a", "b", "c"}
	for k := range x {
		fmt.Println(k) // prints 0, 1, 2
	}

	for _, v := range x {
		fmt.Println(v) // prints a, b, c
	}

	for k, v := range x {
		fmt.Printf("%d => %s\n", k, v)
	}

	// 遍历映射
	var m = map[string]int{
		"go":         1,
		"python":     2,
		"javascript": 3,
	}
	for k, v := range m {
		fmt.Println("Key:", k, ",Value:", v)
	}
}

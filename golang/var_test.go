package golang

import (
	"fmt"
	"testing"
)

var (
	a        = 1
	b uint64 = 3
	c        = 3.14
)

func TestVar(t *testing.T) {
	// 声明
	var v1 *int           //指针
	var v2 map[string]int //key 为string value为int的map
	var v3 func(a int) int
	fmt.Printf("v1=%p, v2=%v, v3=%v\n", v1, v2, v3)

	// 初始化
	var v4 int = 10
	var v5 = 10 //省略类型
	v6 := 10    //此方式只限局部变量
	fmt.Printf("v4=%d, v5=%d, v6=%d\n", v4, v5, v6)

	// 批量定义
	var v7, v8 = 1, 3.14
	fmt.Printf("v7=%d, v8=%f\n", v7, v8)
}

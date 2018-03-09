package base

import (
	"fmt"
	"testing"
)

func TestArray(t *testing.T) {
	// 声明
	var a [3]int

	// 初始化，长度固定
	a = [3]int{1, 2}

	// 赋值
	a[2] = 3
	fmt.Println(a)

	// 省略长度
	b := [...]int{1, 2, 3}
	fmt.Println(len(b))

	// 指定索引
	c := [5]int{2: 50, 3: 100}
	fmt.Println(c)

}

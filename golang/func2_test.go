package base

import (
	"fmt"
	"testing"
)

/**
闭包：
其实理解闭包的最方便的方法就是将闭包函数看成一个类，一个闭包函数调用就是实例化一个类。
然后就可以根据类的角度看出哪些是“全局变量”，哪些是“局部变量”了。
比如上例中的adder函数返回func(int) int 的函数
pos和neg分别实例化了两个“闭包类”，在这个“闭包类”中有个“闭包全局变量”sum。所以这样就很好理解返回的结果了。
 */
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func TestFunc2(t *testing.T) {
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			i,
			pos(i),
			neg(-2*i),
		)
	}
}
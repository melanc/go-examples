package golang

import (
	"fmt"
	"testing"
)

// 声明函数类型
type A func(int, int)

func (f A) Serve() {
	fmt.Println("serve2")
}

func serve(int, int) {
	fmt.Println("serve1")
}

func TestFunc3(t *testing.T) {
	a := A(serve)
	a(1, 2)
	a.Serve()
}

package golang

import (
	"fmt"
	"testing"
)

func test(x int) {
	defer fmt.Println("aaaaaa")

	defer fmt.Println("bbbbbb")

	defer func() {
		fmt.Println("xxxxxx")
	}()

	fmt.Printf("finished: %d\n", x)
}

func test2() {
	fmt.Println("test2")
}

func TestDefer(t *testing.T) {
	test(0)
	test2()
}

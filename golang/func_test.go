package golang

import "fmt"
import (
	"errors"
	"testing"
)

func Add(a int, b int) (ret int, err error) {
	if a < 0 || b < 0 { // 假定只支持两个非负数字的加法
		err = errors.New("Should be non-negative numbers!")
		return
	}
	return a + b, nil
}

func TestFunc(t *testing.T) {
	fmt.Println(Add(-1, 1))
	fmt.Println(Add(10, 1))
}

package base

import (
	"fmt"
	"testing"
)

// 接口
type Handler3 interface {
	Do(k, v interface{})
}

// 实现了Handler接口
type HandlerFunc3 func(k, v interface{})

func (f HandlerFunc3)Do(k, v interface{}){
	f(k, v)
}

func Each3(m map[interface{}]interface{}, h Handler3) {
	if m != nil && len(m) > 0 {
		for k, v := range m {
			h.Do(k, v)
		}
	}
}

func EachFunc(m map[interface{}]interface{}, f func(k, v interface{})) {
	Each3(m, HandlerFunc3(f))
}


func selfInfo(k, v interface{}) {
	fmt.Printf("大家好,我叫%s,今年%d岁\n", k, v)
}

func TestInterface3(t *testing.T) {
	persons := make(map[interface{}]interface{})
	persons["张三"] = 20
	persons["李四"] = 23
	persons["王五"] = 26

	EachFunc(persons, selfInfo)

}

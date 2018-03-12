package base

import (
	"fmt"
	"testing"
)

type Handler2 interface {
	Do(k, v interface{})
}

// 实现了Handler接口
type HandlerFunc func(k, v interface{})

func (f HandlerFunc)Do(k, v interface{}){
	f(k, v)
}

func Each2(m map[interface{}]interface{}, h Handler2) {
	if m != nil && len(m) > 0 {
		for k, v := range m {
			h.Do(k, v)
		}
	}
}

type welcome2 string

func (w welcome2)selfInfo(k, v interface{}){
	fmt.Printf("%s, 我叫%s，今年%d岁\n", w, k, v)
}

func TestInterface2(t *testing.T){
	m := make(map[interface{}]interface{}, 9)
	m["小明"] = 20
	m["小红"] = 21
	m["小龙"] = 22
	var w welcome2 = "大家好"
	Each2(m, HandlerFunc(w.selfInfo))
}
package main

import "fmt"

type Handler interface {
	Do(k, v interface{})
}

func Each(m map[interface{}]interface{}, h Handler) {
	if m != nil && len(m) > 0 {
		for k, v := range m {
			h.Do(k, v)
		}
	}
}

type welcome string

func (w welcome)Do(k, v interface{}){
	fmt.Printf("%s, 我叫%s，今年%d岁\n", w, k, v)
}

func main(){
	m := make(map[interface{}]interface{}, 9)
	m["小明"] = 20
	m["小红"] = 21
	m["小龙"] = 22
	var w welcome = "大家好"
	Each(m, w)
}
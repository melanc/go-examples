package main

import (
	"fmt"
)

type user struct{ name string }

func main() {
	a := make(map[string]int, 100)
	a["age"] = 24
	if v, ok := a["age"]; ok {
		fmt.Printf("age: %d\n", v)
	}
	if v, ok := a["name"]; ok {
		fmt.Printf("name: %d\n", v)
	}

	fmt.Printf("a len: %d\n", len(a))

	a["score"] = 80

	for k, v := range a {
		fmt.Printf("%s=%d\n", k, v)
	}

	m := map[int]user{
		1: {"wang"},
	}
	fmt.Printf("m: %v\n", m)
	tmp := m[1]
	tmp.name = "dong"
	m[1] = tmp
	fmt.Printf("m: %v\n", m)
	//m[1].name = "he"

}

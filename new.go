package main

import "fmt"

type Rect struct {
	x, y float64
	width, height float64
}

func main(){
	// 返回指针
	var b *bool = new(bool)
	*b = true
	fmt.Println(*b)

	var i *int = new(int)
	*i = 100
	fmt.Println(*i)

	var s *string = new(string)
	*s = "golang"
	fmt.Println(*s)

	var p *[]int = new([]int)
	fmt.Println(*p)

	var m *map[string]int = new(map[string]int)
	fmt.Println(*m)


	var r *Rect = new(Rect)
	r.x = 0
	r.y = 0
	r.width = 100
	r.height = 200
	fmt.Println(*r)
}

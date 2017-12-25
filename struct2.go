package main

import "fmt"

type Student struct {
	id int
	name string
}

func (self *Student) Test() {
	fmt.Printf("%p, %v\n", self, self)
}

func (self Student) TestVal() {
	fmt.Printf("%p, %v\n", &self, self)
}

func main() {
	u := Student{1, "Tom"}
	u.Test()
	// 隐式传递 receiver
	mValue := u.Test
	mValue()

	// 显式传递 receiver
	mExpression := (*Student).Test
	mExpression(&u)
	(*Student).TestVal(&u)

}


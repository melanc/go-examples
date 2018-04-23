package base

import (
	"fmt"
	"testing"
)

type Data struct{
	x int
}

// func ValueTest(self Data);
func (self Data) ValueTest() {
	fmt.Printf("Value: %p\n", &self)
}

// func PointerTest(self *Data);
func (self *Data) PointerTest() {
	fmt.Printf("Pointer: %p\n", self)
}
func TestStruct3(t *testing.T) {
	d := Data{}
	p := &d
	fmt.Printf("Data: %p\n", p)
	d.ValueTest() // ValueTest(d)
	d.PointerTest() // PointerTest(&d)
	p.ValueTest() // ValueTest(*p)
	p.PointerTest() // PointerTest(p)
}

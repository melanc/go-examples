package base

import (
	"fmt"
	"testing"
)

type User struct {
	name string
	age int
}

type Manager struct {
	User
	title string
}

func (Manager) Rest(){
	fmt.Println("resting...")
}

func (*Manager) Work(){
	fmt.Println("working...")
}

func TestStruct(t *testing.T) {
	// 创建方式一：T{...}
	u1 := User{
		name: "wang",
		age: 24,
	}
	u1.name = "Wang"
	fmt.Printf("u1: %v\n", u1)
	
	// 创建方式二，指针：&T{...}
	u2 := &User{
		name: "dong",
		age: 30,
	}
	u2.name = "Dong"
	(*u2).age = 28	// 都可以
	fmt.Printf("u2: %v\n", u2)

	// 创建方式三，指针：new(T)
	u3 := new(User)
	u3.name = "tom"
	u3.age = 30
	fmt.Printf("u3: %v\n", u3)

	m := Manager{
		User{"peter", 25},
		"CTO",
	}
	fmt.Printf("m: %v\n", m)

	m.Rest()
	m.Work()
}
package main

import "fmt"

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

func main() {
	u1 := User{
		name: "wang",
		age: 24,
	}
	fmt.Printf("u1: %v\n", u1)

	u1.name = "dong"
	fmt.Printf("u2: %v\n", u1)

	u2 := new(User)
	u2.name = "tom"
	u2.age = 30
	fmt.Printf("u2: %v\n", u2)

	m := Manager{
		User{"peter", 25},
		"CTO",
	}
	fmt.Printf("m: %v\n", m)

	m.Rest()
	m.Work()
}
package main

import "fmt"

func main(){
	s := "acd"
	str := "abcd"
	for i:=0; i<len(str); i++ {
		t := str[0:i] + str[i+1:]
		if t == s {
			fmt.Println(t)
		}
	}
	fmt.Println("no")
}
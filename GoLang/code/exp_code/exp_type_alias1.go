package main

import (
	"fmt"
	"reflect"
)

type User struct {
}
type MyUser1 User
type MyUser2 = User

func (i MyUser1) m1() {
	fmt.Println("MyUser1.m1")
}
func (i User) m2() {
	fmt.Println("User.m2")
}

func main() {
	var i1 MyUser1
	var i2 MyUser2
	i1.m1()
	if reflect.TypeOf(i1) == reflect.TypeOf(i2) {
		fmt.Println("类型相同")
	} else {
		fmt.Println("类型不同")
	}
	//i1.m2()
	i2.m2()
}

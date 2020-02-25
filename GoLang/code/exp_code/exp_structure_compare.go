package main

import "fmt"

func main() {
	sn1 := struct {
		age  int
		name string
	}{age: 11, name1: "qq"}
	sn2 := struct {
		age  int
		name string
	}{age: 11, name: "qq"}
	sn3 := struct {
		name string
		age  int
	}{name: "qq", age: 11}

	if sn1 == sn2 {
		fmt.Println("sn1 == sn2")
	}
	if sn1 == sn3 {
		fmt.Println("sn1 == sn3")
	}
	sm1 := struct {
		age int
		m   map[string]string
	}{age: 11, m: map[string]string{"a": "1"}}
	sm2 := struct {
		age int
		m   map[string]string
	}{age: 11, m: map[string]string{"a": "1"}}
	if sm1 == sm2 {
		fmt.Println("sm1 == sm2")
	}

	sm3 := struct {
		age int
		x chan int
	}{age: 11, x:make(chan int)}
	sm4 := struct {
		age int
		x chan int
	}{age: 11, x:make(chan int)}
	if sm3 == sm4 {
		fmt.Println("sm1 == sm2")
	}
}

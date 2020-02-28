package main

import "fmt"

type Param map[string]interface{}

type Show struct {
	Param
}

type student struct {
	Name string
}

func zhoujielun(v interface{}) {
	switch msg := v.(type) {
	case *student, student:
		fmt.Print(msg.(student).Name)
		//fmt.Println(msg.Name) //error
	}
}

func main() {
	//s := new(Show)
	//s.Param["RMB"] = 10000 //error
}

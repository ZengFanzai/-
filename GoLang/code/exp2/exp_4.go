package main

import (
	"encoding/json"
	"fmt"
)

type Param map[string]interface{}

type Show struct {
	Param
}

type student struct {
	Name string
}

type People struct {
	name string `json:"name"`
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

	js := `{
		"name":"11"
	}`
	var p People
	err := json.Unmarshal([]byte(js), &p)
	if err != nil {
		fmt.Println("err: ", err)
		return
	}
	fmt.Println("people: ", p)
}

package main

import "fmt"

func main() {
	i := GetValue()

	switch i.(type) {
	case int:
		println("int")
		i1, err := i.(int)
		fmt.Println(i1+2, err)
	case string:
		println("string")
	case interface{}:
		println("interface")
	default:
		println("unknown")
	}
}

func GetValue() interface{} {
	//return "string"
	return 1
}
func funcMui(x,y int) int {
	return x+y
}
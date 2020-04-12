package main

import "fmt"

type data struct {
	name string
}

type printer interface {
	print()
}

func (p *data) print() {
	fmt.Println("name: ", p.name)
}

func main() {
	d1 := data{"one"}
	d1.print()    // d1 变量可寻址，可直接调用指针 receiver 的方法

	var in printer = &data{"two"}
	in.print()    // 类型不匹配

	//m := map[string]data{
	//	"x": data{"three"},
	//}
	//m["x"].print()    // m["x"] 是不可寻址的    // 变动频繁
}


package main

var gvar int

func main() {
	//var test int
	one := 0
	one, two := 1, 2 //two是新变量，允许one的重复声明。
	println("hello")
	println(one, two)
	x := 1
	println(x)		// 1
	{
		println(x)	// 1
		x := 2
		println(x)	// 2	// 新的 x 变量的作用域只在代码块内部
	}
	println(x)		// 1

	var m map[string]int
	//m = make(map[string]int)// map 的正确声明，分配了实际的内存
	m["one"] = 1		// error: panic: assignment to entry in nil map
}

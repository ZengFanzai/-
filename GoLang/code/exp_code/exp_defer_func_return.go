package main

import "fmt"

func main() {

	println(DeferFunc1(1))
	println(DeferFunc2(1))
	println(DeferFunc3(1))
	println(DeferFunc4())
	println(DeferFunc5())
	println(DeferFunc6())
}

func DeferFunc1(i int) (t int) {
	t = i
	defer func() {
		t += 3
	}()
	return t
}

//t = 1
//return t ==> return 1
//t+3=4
func DeferFunc2(i int) int {
	t := i
	defer func() {
		t += 3
	}()
	return t
}

func DeferFunc3(i int) (t int) {
	defer func() {
		t += i
	}()
	return 2
}

func DeferFunc4() (r int) {
	t := 5
	defer func() {
		t = t + 5
	}()
	return t
}

// 返回值的r与defer的r不相同
//地址分别为：0xc00000a0c8，0xc00000a100
func DeferFunc5() (r int) {
	fmt.Println(&r)
	defer func(r int) {
		r = r + 5
		fmt.Println(&r)
	}(r)
	return 1
}

//此处两个r地址相同
func DeferFunc6() (r int) {
	fmt.Println(&r)
	defer func() {
		r = r + 5
		fmt.Println(&r)
	}()
	return 1
}

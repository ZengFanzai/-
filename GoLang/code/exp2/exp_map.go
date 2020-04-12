package main

import (
	"fmt"
	"unsafe"
)

type Test struct {
	value int
	nums  []int
}

func main() {
	m := make(map[string]int, 2)
	fmt.Println(unsafe.Pointer(&m))
	m["one"] = 1
	m["two"] = 2
	m["three"] = 3
	fmt.Println(unsafe.Pointer(&m), len(m))
	s := []int{1, 2, 3}
	t := Test{value: 1, nums: []int{1, 2, 3}}
	func(a map[string]int, s []int, t Test) {
		//can
		m["one"] = -1
		//can
		s[0] = 22
		// can't
		t.value = 233
		// can
		t.nums[1] = 334
	}(m, s, t)
	fmt.Println(m, s, t)
}

package main

import "fmt"

func main() {
	s := make([]int, 5)
	s = append(s, 1, 2, 3)
	fmt.Println(s)
	str := make([]string, 5)
	str = append(str, "a", "b", "c")
	fmt.Println(str)
	s1 := make([]int, 0)
	fmt.Println(s1)
	s1 = append(s1, 1, 2, 3)
	fmt.Println(s1)
}

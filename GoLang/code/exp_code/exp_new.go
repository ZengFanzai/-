package main

import "fmt"

func main() {
	s1 := new([]int)
	s2 := []int{1}
	s1 = &s2
	fmt.Println(*s1)
}

/*func main() {
	list:=make([]int,0)
	list = append(list, 1)
	fmt.Println(list)
}*/
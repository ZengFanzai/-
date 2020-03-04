package main

import "fmt"

func outOfRange() int {
	arr := [3]int{1, 2, 3}
	i := 2
	elem := arr[i]
	return elem
}

//func main() {
//	//arr := [3]int{1, 2, 3}
//	////arr1 := [1]int{1}
//	////arr2 := []int{1}
//	//fmt.Println(arr)
//	//outOfRange()
//
//	ch := make(chan int)
//	for i := 1; i < 11; i++ {
//		go func(x int) {
//			ch <- x
//		}(i)
//	}
//	for {
//		select {
//		case v := <-ch:
//			fmt.Print(v, " ")
//		default:
//
//		}
//	}
//}

func main() {
	defer fmt.Println("in main")
	defer func() {
		defer func() {
			defer func() {
				if e := recover(); e!= nil {
					fmt.Println(e)
				}
			}()
			if e := recover(); e!= nil {
				fmt.Println(e)
			}
			panic("panic again and again")
		}()
		if e := recover(); e!= nil {
			fmt.Println(e)
		}
		panic("panic again")
	}()

	panic("panic once")
}

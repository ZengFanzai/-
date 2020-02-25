package main

import "runtime"

//func main() {
//	done := false
//
//	go func() {
//		done = true
//	}()
//
//	for !done {
//	}
//
//	println("done !")
//}

//func main() {
//	done := false
//
//	go func() {
//		done = true
//	}()
//
//	for !done {
//		println("not done !")	// 并不内联执行
//	}
//
//	println("done !")
//}

func main() {
	done := false

	go func() {
		done = true
	}()

	for !done {
		println("test")
		runtime.Gosched()
	}

	println("done !")
}
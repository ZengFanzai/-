package main

import (
	"fmt"
	"time"
)

//func main() {
//	ch := make(chan string)
//
//	go func() {
//		for m := range ch {
//			fmt.Println("Processed:", m)
//			time.Sleep(1 * time.Second) // 模拟需要长时间运行的操作
//		}
//	}()
//
//	ch <- "cmd.1"
//	ch <- "cmd.2" // 不会被接收处理
//}

//func main() {
//	ch := make(chan int)
//	for i := 0; i < 3; i++ {
//		go func(idx int) {
//			ch <- idx
//		}(i)
//	}
//
//	fmt.Println(<-ch)		// 输出第一个发送的值
//	close(ch)			// 不能关闭，还有其他的 sender
//	time.Sleep(2 * time.Second)	// 模拟做其他的操作
//}

//func main() {
//	ch := make(chan int)
//	done := make(chan struct{})
//
//	for i := 0; i < 3; i++ {
//		go func(idx int) {
//			select {
//			case ch <- (idx + 1) * 2:
//				fmt.Println(idx, "Send result")
//			case <-done:
//				fmt.Println(idx, "Exiting")
//			}
//		}(i)
//	}
//
//	fmt.Println("Result: ", <-ch)
//	close(done)
//	time.Sleep(3 * time.Second)
//}

//func main() {
//	var ch chan int // 未初始化，值为 nil
//	for i := 0; i < 3; i++ {
//		go func(i int) {
//			ch <- i
//		}(i)
//	}
//
//	fmt.Println("Result: ", <-ch)
//	time.Sleep(2 * time.Second)
//}

func main() {
	inCh := make(chan int)
	outCh := make(chan int)

	go func() {
		var in <-chan int = inCh
		var out chan<- int
		var val int

		for {
			select {
			case out <- val:
				println("--------")
				out = nil
				in = inCh
			case val = <-in:
				println("++++++++++")
				out = outCh
				in = nil
			}
		}
	}()

	go func() {
		for r := range outCh {
			fmt.Println("Result: ", r)
		}
	}()

	time.Sleep(0)
	inCh <- 1
	inCh <- 2
	time.Sleep(3 * time.Second)
}
package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		// 1 在这里需要你写算法
		// 2 要求每秒钟调用一次proc函数
		// 3 要求程序不能退出
		x := func() {
			defer func() {
				if err := recover(); err != nil {
					fmt.Println(err)
				}
			}()
			proc()
		}
		for {
			select {
			case <-time.Tick(1 * time.Second):
				x()
			}
		}
	}()

	select {}
}

func proc() {
	panic("ok")
}

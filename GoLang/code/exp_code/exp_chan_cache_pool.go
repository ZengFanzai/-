package main

import (
	"fmt"
	"sync"
)

//下面的迭代会有什么问题？

type threadSafeSet struct {
	sync.WaitGroup
	s []interface{}
}

func (set *threadSafeSet) Iter() <-chan interface{} {
	ch := make(chan interface{})
	defer set.WaitGroup.Wait()
	//ch := make(chan interface{}, )
	for elem := range set.s {
		set.WaitGroup.Add(1)
		//ch <- elem
		go func() {
			defer set.WaitGroup.Done()
			ch <- elem
		}()
		//close(ch)
	}
	//go func() {
	//	for elem := range set.s {
	//		ch <- elem
	//	}
	//	close(ch)
	//}()
	return ch
}

func main() {
	th := threadSafeSet{
		s: []interface{}{"1", "2", "3","1", "2", "3","1", "2", "3","1", "2", "3"},
	}
	go func() {
		for i := range th.Iter() {
			fmt.Println(i)
		}
	}()

}

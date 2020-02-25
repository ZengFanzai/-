package main

import (
	"fmt"
	"sync"
)

type UserAges struct {
	ages map[string]int
	sync.Mutex
}

func (ua *UserAges) Add(name string, age int) {
	ua.Lock()
	defer ua.Unlock()
	ua.ages = make(map[string]int)
	ua.ages[name] = age
	//or
	//ua.ages = map[string]int{name:age}
}

func (ua *UserAges) Get(name string) int {
	if age, ok := ua.ages[name]; ok {
		return age
	}
	return -1
}

func (ua *UserAges) Test(chan1 chan int, name string)  {
	age := ua.Get(name)
	chan1 <- age
}

func main() {
	ua := new(UserAges)
	chan1 := make(chan int)
	ua.Add("test", 20)
	go ua.Test(chan1, "test")
	fmt.Println(<-chan1)
}
package main

import (
	"encoding/json"
	"fmt"
	"sync"
)

type Stduent struct {
	Age int
}

type Girl struct {
	Name       string `json:"name"`
	DressColor string `json:"dress_color"`
}

func (g Girl) SetColor(color string) {
	g.DressColor = color
}
func (g Girl) JSON() string {
	data, _ := json.Marshal(&g)
	return string(data)
}

type query func(string) string

func exec(wg *sync.WaitGroup, name string, vs ...query) string {
	ch := make(chan string)
	s := ""
	fn := func(i int) {
		ch <- vs[i](name)
		wg.Done()
	}
	for i, _ := range vs {
		wg.Add(1)
		go fn(i)
		s += <-ch
	}
	return s
}

func main() {
	kv := map[string]Stduent{"menglu": {Age: 21}}
	//kv["menglu"].Age = 22 //error
	s := []Stduent{{Age: 21}}
	s[0].Age = 22
	fmt.Println(kv, s)

	fmt.Println([...]string{"1", "2"} == [...]string{"1", "1"})
	//fmt.Println([]string{"1"} == []string{"1"}) //error
	fmt.Println(Stduent{Age: 1} == Stduent{Age: 1})

	str1 := []string{"a", "b", "c"}
	str2 := str1[1:]
	str2[1] = "new"
	fmt.Println(str1)
	str2 = append(str2, "z", "x", "y")
	fmt.Println(str1)

	g := Girl{Name: "menglu"}
	g.SetColor("white")
	fmt.Println(g.JSON())

	wg := &sync.WaitGroup{}
	ret := exec(wg, "111", func(n string) string {
		return n + "func1"
	}, func(n string) string {
		return n + "func2"
	}, func(n string) string {
		return n + "func3"
	}, func(n string) string {
		return n + "func4"
	})
	wg.Wait()
	fmt.Println(ret)

	const (
		a = iota //0
		b = iota //1
	)
	const (
		name = "menglu" //0
		c    = iota //1
		d    = iota //2
	)

	fmt.Println(a,b,c,d)
}

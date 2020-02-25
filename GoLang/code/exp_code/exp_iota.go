package main

import "fmt"

const (
	x = iota
	y
	z = "zz"
	k
	p = iota
)

const (
	i = 1
	j = iota
	k1
	l
)

const (
	Apple, Banana = iota + 1, iota + 2
	Cherimoya, Durian
	Elderberry, Fig
)


type AudioOutput int

const (
	i1 = iota
	j1
	_
	_
	k2
	l1
)

func main() {
	fmt.Println(x, y, z, k, p)
	fmt.Println(i, j, k1, l)
	fmt.Println(Apple, Banana,"\n", Cherimoya, Durian,"\n", Elderberry, Fig)
	fmt.Println(i1,j1,k2,l1)
}

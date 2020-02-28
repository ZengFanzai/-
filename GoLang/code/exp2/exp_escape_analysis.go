package main

type Test struct {
	name string
}

func getTest() (*Test, *[]int) {
	i := make([]int, 10)
	i[1] = 1
	test := Test{name: "dasdd"}
	return &test, &i
}

func main() {
	//s := "Escape"
	//fmt.Println(s)
	getTest()
}

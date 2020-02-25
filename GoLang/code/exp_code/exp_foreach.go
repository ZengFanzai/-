package main

import "fmt"

type student struct {
	Name string
	Age  int
}

func pase_student() map[string]*student {
	m := make(map[string]*student)
	stus := []student{
		{Name: "zhou", Age: 24},
		{Name: "li", Age: 23},
		{Name: "wang", Age: 22},
	}

	//以m[stu.Name]=&stu实际上一致指向同一个指针， 最终该指针的值为遍历的最后一个struct的值拷贝
	//错误语法
	for _, stu := range stus {
		fmt.Println(stus, &stu)
		m[stu.Name] = &stu
	}
	fmt.Println(m)

	for i := 0; i < len(stus); i++ {
		m[stus[i].Name] = &stus[i]
	}
	fmt.Println(m)
	return m

}

func main() {
	pase_student()
}

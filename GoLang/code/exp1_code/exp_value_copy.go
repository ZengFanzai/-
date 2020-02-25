package main

import "fmt"

// 数组使用值拷贝传参
func main() {
	x := [3]int{1,2,3}

	func(arr [3]int) {
		arr[0] = 7
		fmt.Println(arr)	// [7 2 3]
	}(x)
	fmt.Println(x)			// [1 2 3]	// 并不是你以为的 [7 2 3]

	func(arr *[3]int) {
		arr[0] = 7
		fmt.Println(arr)	// [7 2 3]
	}(&x)
	fmt.Println(x)			// [7 2 3]

	x2 := x[:]
	func(arr []int) {
		arr[0] = 7
		fmt.Println(x)	// [7 2 3]
	}(x2)
	fmt.Println(x)	// [7 2 3]

}

func pointerUse(arr *[3]int) {
	(*arr)[0] = 7
	fmt.Println(*arr) // [7 2 3]
}

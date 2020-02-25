package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(reverse(1234567899))
	fmt.Println(reverse(12345))
	fmt.Println(reverse(-123))
}

func reverse(x int32) int32 {
	var num int64
	x64 := int64(x)
	for x64 != 0 {
		num = num*10 + x64%10
		x64 = x64 / 10
	}
	// 使用math包中的最大最小值
	if num > math.MaxInt32 || num < math.MinInt32 {
		return 0
	}
	return int32(num)
}

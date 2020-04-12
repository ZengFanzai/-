package main

import "fmt"

var bitArr []int

func lowbit(x int) int {
	return x & -x
}

func build(list []int) {
	bitArr = make([]int, len(list)+1)
	for i := 1; i <= len(list); i++ {
		for j := i - lowbit(i); j <= i-1; j++ {
			bitArr[i] += list[j]
		}
	}
}

func add(idx int, delta int) {
	//idx代表原数组的下标，使用需要加1跳过bitArr[0]
	for i := idx+1; i < len(bitArr); i += lowbit(i) {
		bitArr[i] += delta
	}
}

func sum(idx int) int {
	result := 0
	for i := idx+1; i > 0; i -= lowbit(i) {
		result += bitArr[i]
	}
	return result
}

func rangeSum(fromIdx, toIdx int) int {
	return sum(toIdx) - sum(fromIdx-1)
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}
	build(arr)
	fmt.Println(sum(4))
	fmt.Println(bitArr)
	add(13, 1)
	fmt.Println(bitArr)
	//add(6, 1)
	//fmt.Println(bitArr)
	//fmt.Println(sum(0))
	//fmt.Println(sum(1))
	//fmt.Println(rangeSum(1,2))
}

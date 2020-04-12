package main

import (
	"fmt"
	"strconv"
)

func test(n int) []string {
	ret := make([]string, n)
	for i := 1; i <= n; i++ {
		if i%15 == 0 {
			ret[i-1] = "FizzBuzz"
		} else if i%3 == 0 {
			ret[i-1] = "Fizz"
		} else if i%5 == 0 {
			ret[i-1] = "Buzz"
		} else {
			ret[i-1] = strconv.Itoa(i)
		}
	}
	return ret
}

func test1(nums []int) int {
	mp := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		mp[nums[i]] += 1
	}
	for k, v := range mp {
		if v == 1 {
			return k
		}
	}
	// 此次在当前题目中是无效的，因为一定存在v==1
	return -1
}

func main() {
	fmt.Println(test(15))
	fmt.Println(test1([]int{1,2,2,1,3,4,5,3,4}))
}

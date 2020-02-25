package main

import (
	"fmt"
	"sort"
)

type Interval struct {
	Start int
	End   int
}

func merge(intervals []Interval) []Interval {
	if len(intervals) <= 1 {
		return intervals
	}

	// 先按start的大小排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i].Start < intervals[j].Start
	})

	res := make([]Interval, 0)
	swap := Interval{}
	for k, v := range intervals {
		if k == 0 {
			swap = v
			continue
		}
		if v.Start <= swap.End {
			swap.End = v.End
		} else {
			res = append(res, swap)
			swap = v
		}
	}
	res = append(res, swap)
	return res
}

func main() {
	intervals := []Interval{
		{Start:1, End:3},
		{Start:8, End:10},
		{Start:2, End:6},
		{Start:15, End:18},
	}
	fmt.Println(merge(intervals))
}
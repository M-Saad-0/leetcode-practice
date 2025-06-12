package main

import (
	"sort"
)

func maximumValueSum(nums []int, k int, edges [][]int) int64 {
	// var num int64 = 0
	delta := make([]int64, len(nums))
	for i := 0; i < len(nums); i++ {
		delta[i] = int64(nums[i]^k) - int64(nums[i])
	}
	sort.Slice(delta, func(i, j int) bool {
		return delta[i] > delta[j]
	})
	total := int64(0)
	for _, v := range nums {
		total += int64(v)
	}
	for i := 0; i < len(delta); i += 2 {
		if i+1 >= len(delta) {
			continue
		}
		a, b := delta[i], delta[i+1]
		if a+b > 0 {
			total += a + b
		}
	}
	return total
}

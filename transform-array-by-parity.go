package main

import (
	"sort"
)

func transformArray(nums []int) []int {

	for i, _ := range nums {
		nums[i] &= 1
	}
	sort.Sort((sort.IntSlice(nums)))
	return nums
}


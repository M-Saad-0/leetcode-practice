package main

func singleNumber(nums []int) int {
	c := 0
	for _, v := range nums {
		c ^= v
	}
	return c
}

package main

import "fmt"

func main() {
	nums := []int{1, 1, 2, 2}
	n := 2
	fmt.Println(shuffle(nums, n))
}

func shuffle(nums []int, n int) []int {
	result := make([]int, 2*n)
	for i := 0; i < n; i++ {
		result[i*2] = nums[i]
		result[i*2+1] = nums[n+i]
	}
	return result
}

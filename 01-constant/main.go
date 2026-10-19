package main

import "fmt"

func getFirst(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	return nums[0]
}

func main() {
	nums := []int{10, 20, 30, 40, 50}

	result := getFirst(nums)
	fmt.Println("First element:", result)
}

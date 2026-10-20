package main

import "fmt"

func linearSearch(nums []int, target int) int {
	for i, num := range nums {
		if num == target {
			return i
		}
	}

	return -1
}

func main() {
	nums := []int{10, 20, 30, 40, 50, 60, 70, 80}
	target := 60

	index := linearSearch(nums, target)
	fmt.Println("Target index:", index)
}

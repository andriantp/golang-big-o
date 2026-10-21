package main

import "fmt"

func comparePairs(nums []int) {
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			fmt.Printf("(%d, %d)\n", nums[i], nums[j])
		}
	}
}

func main() {
	nums := []int{10, 20, 30}

	comparePairs(nums)
}

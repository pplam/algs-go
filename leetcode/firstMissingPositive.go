package main

import "fmt"

func main() {
	nums := []int{1,2,0}
	smm := firstMissingPositive(nums)
	fmt.Println(smm)
}

func firstMissingPositive(nums []int) int {
	smallestMissing := 1
	for i := 0; i < len(nums); i++ {
		if nums[i] >= 0 && (nums[i] == smallestMissing || nums[i] + 1 < smallestMissing) {
			smallestMissing = nums[i] + 1
		}
	}
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] >= 0 && (nums[i] == smallestMissing || nums[i] + 1 < smallestMissing) {
			smallestMissing = nums[i] + 1
		}
	}
	return smallestMissing
}

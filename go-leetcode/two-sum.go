package main

import "fmt"

func twoSum(nums []int, target int) []int {
	res := make([]int, 2)
	res[0] = -1
	res[1] = -1
	for i := 0; i < len(nums)-1; i++ {
		for j := i; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				res[0] = i
				res[1] = j
			}
		}
	}

	return res
}

var pow = []int{1, 2, 4, 8, 16, 32, 64, 129}

func main() {
	x := twoSum(pow, 130)
	fmt.Println(x)
}

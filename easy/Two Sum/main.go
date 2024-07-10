package main

import "fmt"

//	Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.
//
//	You may assume that each input would have exactly one solution, and you may not use the same element twice.
//
//	You can return the answer in any order.
//
//
//
//	Example 1:
//
//	Input: nums = [2,7,11,15], target = 9
//	Output: [0,1]
//	Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].
//	Example 2:
//
//	Input: nums = [3,2,4], target = 6
//	Output: [1,2]
//	Example 3:
//
//	Input: nums = [3,3], target = 6
//	Output: [0,1]

// First O(n*n) solution
var (
	nums      = []int{3, 3}
	targetSum = 6
	output    []int
)

//func main() {
//	for i, n := range nums {
//		for targetIndex := i + 1; targetIndex < len(nums); targetIndex++ {
//			if n+nums[targetIndex] == targetSum {
//				output = append(output, i, targetIndex)
//			}
//		}
//	}
//	fmt.Println(output)
//}

// Second way with O(n)
func main() {
	ourHashMap := make(map[int]int)

	for index, number := range nums {
		cal := targetSum - number

		valueIndex, exits := ourHashMap[cal]
		if exits {
			output = append(output, index, valueIndex)
		} else {
			ourHashMap[number] = index
		}

	}
	fmt.Println(output)
}

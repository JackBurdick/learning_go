/*
from leetcode:
	"Given an array of integers, return indices of the two numbers such that
	they add up to a specific target. You may assume that each input would have
	exactly one solution, and you may not use the same element twice."

Example solution:
	Given;
		- nums = [2, 7, 11, 15]
		- target = 9
	Solution;
		- [0, 1]

Rational;
	nums[0] + nums[1] => 2 + 7 = 9,
	therefore; return [0, 1]
*/

package twoSum

import "fmt"

// func twoSum(nums []int, target int) ([]int, error) {
// 	// brute force
//	// Time Complexity: O(n^2)
//	// Space complexity: O(1)
// 	for i, valA := range nums {
// 		for j, valB := range nums[i+1:] {
// 			if valA+valB == target {
// 				// offset second index by i + 1
// 				return []int{i, j + i + 1}, nil
// 			}
// 		}
// 	}
// 	return nil, fmt.Errorf("No values summed to %v", target)
// }

// func twoSum(nums []int, target int) ([]int, error) {
// 	// Use map to improve performance.
// 	// Time complexity is O(n) since we traverse the list exactly twice.
// 	// Space Complexity is O(n) since we store the map.
// 	// benefit here is the improved time complexity, but the cost is the
// 	// additional space complexity
// 	m := make(map[int]int)

// 	for i, v := range nums {
// 		// populate map
// 		m[v] = i
// 	}
// 	for i, v := range nums {
// 		// use map to find complement index
// 		comp := target - v
// 		j, ok := m[comp]
// 		if ok {
// 			return []int{i, j}, nil
// 		}
// 	}
// 	return nil, fmt.Errorf("No values summed to %v", target)
// }

func twoSum(nums []int, target int) ([]int, error) {
	// Use map to improve performance, but find the solution in only one pass.
	// Time complexity is O(n) since we traverse the list once.
	// Space Complexity is O(n) since we store the map.
	// The benefit here is that we only make one pass.
	m := make(map[int]int)
	for i, v := range nums {
		comp := target - v
		j, ok := m[comp]
		switch ok {
		case true:
			return []int{j, i}, nil
		default:
			m[v] = i
		}
	}
	return nil, fmt.Errorf("No values summed to %v", target)
}

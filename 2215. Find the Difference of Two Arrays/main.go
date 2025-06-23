package main

import "fmt"

func contains(nums []int, search int) bool {
	for _, el := range nums {
		if el == search {
			return true
		}
	}
	return false
}

func findDifference(nums1 []int, nums2 []int) [][]int {
	result := [][]int{{}, {}}
	for _, num1 := range nums1 {
		if !contains(nums2, num1) && !contains(result[0], num1) {
			result[0] = append(result[0], num1)
		}
	}

	for _, num2 := range nums2 {
		if !contains(nums1, num2) && !contains(result[1], num2) {
			result[1] = append(result[1], num2)
		}
	}
	return result
}

func main() {
	nums1 := []int{1, 2, 3}
	nums2 := []int{2, 4, 6}
	fmt.Println(findDifference(nums1, nums2))
}

package main

func pivotIndex(nums []int) int {
	sumr := 0
	total := 0
	suml := 0

	for i := range nums {
		total += nums[i]
	}
	for i := range nums {
		sumr = total - suml - nums[i]
		if suml == sumr {
			return i
		}
		suml += nums[i]
	}

	return -1
}

func main() {
	nums := []int{1, 7, 3, 6, 5, 6}
	pivotIndex(nums)
}

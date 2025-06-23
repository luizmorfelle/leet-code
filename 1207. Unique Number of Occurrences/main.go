package main

import (
	"slices"
	"strconv"
)

func uniqueOccurrences(arr []int) bool {
	var mapNumbers = make(map[string]int)

	for _, el := range arr {
		actual := mapNumbers[strconv.Itoa(el)]

		if actual == 0 {
			mapNumbers[strconv.Itoa(el)] = 1
		} else {
			mapNumbers[strconv.Itoa(el)] = actual + 1
		}
	}

	visited := []int{}

	for _, value := range mapNumbers {
		if slices.Contains(visited, value) {
			return false
		} else {
			visited = append(visited, value)
		}
	}
	return true
}

func main() {
	numbers := []int{7, -3, -3, 7, 7, -2}
	uniqueOccurrences(numbers)
}

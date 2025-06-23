package main

import "fmt"

func largestAltitude(gain []int) int {
	largest := 0
	sum := 0

	for _, alt := range gain {
		sum += alt

		if sum > largest {
			largest = sum
		}
	}
	return largest
}

func main() {
	gains := []int{-4, -3, -2, -1, 4, 3, 2}
	fmt.Println(largestAltitude(gains))
}

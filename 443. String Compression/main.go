package main

import (
	"fmt"
	"strconv"
)

func compress(chars []byte) int {
	newChars := ""
	count := 0
	var control byte
	for i := range chars {
		actual := chars[i]
		if actual != control {
			control = actual
			if count > 1 {
				newChars += strconv.Itoa(count)
			}
			newChars += string(actual)
			count = 0
		}
		count++
	}
	if count > 1 {
		newChars += strconv.Itoa(count)
	}
	chars = []byte(newChars)
	return len(newChars)
}

func main() {
	chars := []byte{'a', 'a', 'b', 'b', 'c', 'c', 'c'}
	fmt.Println("RESULT:", compress(chars))
}

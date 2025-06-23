package main

import "fmt"

func isSubsequence(s string, t string) bool {
	count := 0
	if len(s) == 0 {
		return true
	}
	for _, letter := range []byte(t) {
		if s[count] == letter {
			count++
		}
		if count == len(s) {
			break
		}
	}
	return count == len(s)
}

func main() {
	fmt.Println(isSubsequence("axc", "ahbgdc"))
}

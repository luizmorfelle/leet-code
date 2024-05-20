package main

func lengthOfLongestSubstring(s string) int {

	if len(s) == 0 {
		return 0
	}

	m := make(map[byte]int)
	left, right := 0, 0
	max := 0

	for right < len(s) {
		if _, ok := m[s[right]]; ok {
			left = maxInt(m[s[right]]+1, left)
		}
		m[s[right]] = right
		max = maxInt(max, right-left+1)

		right++
	}

	return max
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	println(lengthOfLongestSubstring("pwwkewwkeabcdedfghijklmnopqrstuvwxyza"))
}
